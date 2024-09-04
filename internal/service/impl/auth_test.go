package impl

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/cripplemymind9/go-market/internal/entity"
	"github.com/cripplemymind9/go-market/internal/mocks/repomocks"
	"github.com/cripplemymind9/go-market/internal/repository/repoerrs"
	"github.com/cripplemymind9/go-market/internal/service/serviceerrs"
	"github.com/cripplemymind9/go-market/internal/service/types"
	"github.com/cripplemymind9/go-market/pkg/hasher"
	"github.com/golang/mock/gomock"
	// "golang.org/x/crypto/bcrypt"
)

func TestAuthService_RegisterUser(t *testing.T) {
	type args struct {
		ctx   context.Context
		input types.AuthRegisterUserInput
	}

	type MockBehaviour func(m *repomocks.MockUser, args args)

	testCases := []struct {
		name          string
		args          args
		mockBehaviour MockBehaviour
		want          int
		wantErr       bool
	}{
		{
			name: "OK",
			args: args{
				ctx: context.Background(),
				input: types.AuthRegisterUserInput{
					Username: "test",
					Password: "Qwerty1!",
					Email:    "test@gmail.com",
				},
			},
			mockBehaviour: func(m *repomocks.MockUser, args args) {
				m.EXPECT().RegisterUser(args.ctx, gomock.Any()).Return(1, nil)
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "User already exists",
			args: args{
				ctx: context.Background(),
				input: types.AuthRegisterUserInput{
					Username: "existingUser",
					Password: "Qwerty2!",
					Email:    "existing@gmail.com",
				},
			},
			mockBehaviour: func(m *repomocks.MockUser, args args) {
				m.EXPECT().RegisterUser(args.ctx, gomock.Any()).Return(0, repoerrs.ErrAlreadyExists)
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Password Hashing Failed",
			args: args{
				ctx: context.Background(),
				input: types.AuthRegisterUserInput{
					Username: "hashFail",
					Password: "123",
					Email:    "hashfail@gmail.com",
				},
			},
			mockBehaviour: func(m *repomocks.MockUser, args args) {
				m.EXPECT().RegisterUser(args.ctx, gomock.Any()).Return(0, serviceerrs.ErrPasswordHashingFailed)
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Cannot Create User",
			args: args{
				ctx: context.Background(),
				input: types.AuthRegisterUserInput{
					Username: "failUser",
					Password: "Fail123!",
					Email:    "fail@gmail.com",
				},
			},
			mockBehaviour: func(m *repomocks.MockUser, args args) {
				m.EXPECT().RegisterUser(args.ctx, gomock.Any()).Return(0, errors.New("unexpected error"))
			},
			want:    0,
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userRepo := repomocks.NewMockUser(ctrl)
			tc.mockBehaviour(userRepo, tc.args)

			passwordHasher := hasher.NewBcryptHasher()
			signKey := "secretkey"
			tokenTTL := time.Hour * 3

			s := NewAuthService(userRepo, passwordHasher, signKey, tokenTTL)
			got, err := s.RegisterUser(tc.args.ctx, tc.args.input)

			if (err != nil) != tc.wantErr {
				t.Errorf("RegisterUser() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if got != tc.want {
				t.Errorf("RegisterUser() = %v, want %v", got, tc.want)
			}
		})
	}
}

var (
	hashedPassword = "$eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjU0NjE2MjAsImlhdCI6MTcyNTQ1MDgyMCwiVXNlcklEIjoxLCJVc2VybmFtZSI6InRlc3QifQ.-DIqnsLMzeueD9MvPv5JpDBj_e6VxiH6_eVVChn8Af0"
)

func TestAuthService_GenerateToken(t *testing.T) {
	type args struct {
		ctx   context.Context
		input types.AuthGenerateTokenInput
	}

	type MockBehaviour func(m *repomocks.MockUser, args args)

	testCases := []struct {
		name          string
		args          args
		mockBehaviour MockBehaviour
		want          string
		wantErr       bool
	}{
		{
			name: "OK",
			args: args{
				ctx: context.Background(),
				input: types.AuthGenerateTokenInput{
					Username: "test",
					Password: "Qwerty1!",
				},
			},
			mockBehaviour: func(m *repomocks.MockUser, args args) {
				// hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(args.input.Password), bcrypt.DefaultCost)
				m.EXPECT().LoginUser(args.ctx, gomock.Any()).Return(entity.User{
					ID:       1,
					Username: "test",
					Password: hashedPassword,
				}, nil)
			},
			want: hashedPassword,
			wantErr: false,
		},
		{
			name: "User Not Found",
			args: args{
				ctx: context.Background(),
				input: types.AuthGenerateTokenInput{
					Username: "unknownuser",
					Password: "Qwerty2!",
				},
			},
			mockBehaviour: func(m *repomocks.MockUser, args args) {
				m.EXPECT().LoginUser(args.ctx, gomock.Any()).Return(entity.User{}, repoerrs.ErrNotFound)
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Invalid Password",
			args: args{
				ctx: context.Background(),
				input: types.AuthGenerateTokenInput{
					Username: "testuser",
					Password: "wrongpassword",
				},
			},
			mockBehaviour: func(m *repomocks.MockUser, args args) {
				m.EXPECT().LoginUser(args.ctx, gomock.Any()).Return(entity.User{
					ID:       1,
					Username: "testuser",
					Password: "hashedPassword",
				}, nil)
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Cannot Sign Token",
			args: args{
				ctx: context.Background(),
				input: types.AuthGenerateTokenInput{
					Username: "testuser",
					Password: "Qwerty3!",
				},
			},
			mockBehaviour: func(m *repomocks.MockUser, args args) {
				m.EXPECT().LoginUser(args.ctx, gomock.Any()).Return(entity.User{
					ID:       1,
					Username: "testuser",
					Password: "hashedPassword",
				}, nil)
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userRepo := repomocks.NewMockUser(ctrl)
			tc.mockBehaviour(userRepo, tc.args)

			passwordHasher := hasher.NewBcryptHasher()
			signKey := "secretkey"
			tokenTTL := time.Hour * 3

			s := NewAuthService(userRepo, passwordHasher, signKey, tokenTTL)

			got, err := s.GenerateToken(tc.args.ctx, tc.args.input)

			if (err != nil) != tc.wantErr {
				t.Errorf("GenerateToken() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !tc.wantErr && got != tc.want {
				t.Errorf("GenerateToken() = %v, want %v", got, tc.want)
			}
		})
	}
}