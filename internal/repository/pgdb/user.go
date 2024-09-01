package pgdb

import (
	"github.com/cripplemymind9/go-market/internal/repository/repoerrs"
	"github.com/cripplemymind9/go-market/internal/entity"
	"github.com/cripplemymind9/go-market/pkg/postgres"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5"
	"context"
	"errors"
	"fmt"
)

type UserRepo struct {
	*postgres.Postgres
}

func NewUserRepo(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

func (r *UserRepo) RegisterUser(ctx context.Context, user entity.User) (int, error) {
	sql, args, err := r.Builder.
		Insert("users").
		Columns("username", "password", "email").
		Values(
			user.Username,
			user.Password,
			user.Email,
		).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("UserRepo.RegisterUser - r.Builder.Insert: %v", err)
	}

	var id int
	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok {
			if pgErr.Code == "23505" {
				return 0, repoerrs.ErrAlreadyExists
			}
		}
		return 0, fmt.Errorf("UserRepo.CreateUser - r.Pool.QueryRow: %v", err)
	}

	return id, nil
}

func (r *UserRepo) LoginUser(ctx context.Context, username, password string) (entity.User, error) {
	sql, args, err := r.Builder.
		Select("*").From("users").
		Where(
			squirrel.Eq{"username": username},
			squirrel.Eq{"password": password},
		).
		ToSql()
	if err != nil {
		return entity.User{}, fmt.Errorf("UserRepo.LoginUser - r.Bulder.Select: %v", err)
	}

	var user entity.User
	err = r.Pool.QueryRow(ctx, sql, args...).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return entity.User{}, fmt.Errorf("user not found")
		}
		return entity.User{}, fmt.Errorf("UserRepo.LoginUser - r.Pool.QueryRow: %v", err)
	}

	return user, nil
}

func (r *UserRepo) GetUserProfile(ctx context.Context, id int) (entity.User, error) {
	sql, args, err := r.Builder.
		Select("*").
		From("users").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return entity.User{}, fmt.Errorf("UserRepo.GetUserProfile - r.Builder.Select: %v", err)
	}

	var user entity.User
	err = r.Pool.QueryRow(ctx, sql, args...).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.User{}, repoerrs.ErrNotFound
		}
		return entity.User{}, fmt.Errorf("UserRepo.GetUserProfile - r.Pool.QueryRow: %v", err)
	}

	return user, nil
}