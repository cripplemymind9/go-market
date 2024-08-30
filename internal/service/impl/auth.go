package impl

import (
	"github.com/cripplemymind9/go-market/internal/repository/repoerrs"
	"github.com/cripplemymind9/go-market/internal/service/serviceerrs"
	"github.com/cripplemymind9/go-market/internal/service/types"
	"github.com/cripplemymind9/go-market/internal/repository"
	"github.com/cripplemymind9/go-market/internal/entity"
	"github.com/cripplemymind9/go-market/pkg/hasher"
	"github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"
	"context"
	"errors"
	"time"
	"fmt"
)

type TokenClaims struct {
	jwt.RegisteredClaims
	UserID 			int
	Username 	string
}

type AuthService struct {
	userRepo 		repository.User
	passwordHasher 	hasher.PasswordHasher
	signKey 		string
	tokenTTL		time.Duration
}

func NewAuthService(userRepo repository.User, passwordHasher hasher.PasswordHasher, signKey string, tokenTTL time.Duration) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		passwordHasher: passwordHasher,
		signKey: signKey,
		tokenTTL: tokenTTL,
	}
}

func (s *AuthService) RegisterUser(ctx context.Context, input types.AuthRegisterUserInput) (int, error) {
	hashedPassword, err := s.passwordHasher.HashPassword(input.Password)
	if err != nil {
		return 0, serviceerrs.ErrPasswordHashingFailed
	}

	user := entity.User{
		Username: input.Username,
		Password: hashedPassword,
		Email: input.Email,
	}

	userId, err := s.userRepo.RegisterUser(ctx, user)
	if err != nil {
		if errors.Is(err, repoerrs.ErrAlreadyExists) {
			return 0, repoerrs.ErrAlreadyExists
		}
		log.Errorf("AuthService.CreateUser - s.userRepo.RegisterUser: %v", err)
		return 0, serviceerrs.ErrCannotCreateUser
	}

	return userId, nil
}

func (s *AuthService) GenerateToken(ctx context.Context, input types.AuthGenerateTokenInput) (string, error) {
	hashedPassword, err := s.passwordHasher.HashPassword(input.Password)
	if err != nil {
		return "", serviceerrs.ErrPasswordHashingFailed
	}
	
	user, err := s.userRepo.LoginUser(ctx, input.Username, hashedPassword)
	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return "", serviceerrs.ErrUserNotFound
		}
		log.Errorf("AuthService.GenerateToken - s.userRepo.LoginUser: %v", err)
		return "", serviceerrs.ErrCannotGetUser
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.tokenTTL)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserID: user.ID,
		Username: user.Username,
	})

	tokenString, err := token.SignedString([]byte(s.signKey))
	if err != nil {
		log.Errorf("AuthService.GenerateToken: caanot sign token: %v", err)
		return "", serviceerrs.ErrCannotSignToken
	}

	return tokenString, nil
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	claims := &TokenClaims{}
	token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil,  fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.signKey), nil
	})
	if err != nil {
		return 0, serviceerrs.ErrCannotParseToken
	}

	if !token.Valid {
		return 0, serviceerrs.ErrCannotParseToken
	}

	return claims.UserID, nil
}