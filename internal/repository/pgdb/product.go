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

type ProductRepo struct {
	*postgres.Postgres
}

func NewProductRepo(pg *postgres.Postgres) *ProductRepo {
	return &ProductRepo{pg}
}

func (r *ProductRepo) CreateProduct(ctx context.Context, product entity.Product) (int, error) {
	sql, args, err := squirrel.
		Insert("products").
		Columns("name", "description", "price", "quantity").
		Values(
			product.Name,
			product.Description,
			product.Price,
			product.Quantity,
		).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("ProductRepo.CreateProduct - squirrel.Insert: %v", err)
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
		return 0, fmt.Errorf("ProductRepo.CreateProduct - r.Pool.QueryRow: %v", err)
	}

	return id, nil
}

func (r *ProductRepo) GetAllProducts(ctx context.Context) ([]entity.Product, error) {
	sql, args, err := squirrel.
		Select("*").
		From("products").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("ProductRepo.GetAllProducts - squirrel.Select: %v", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("ProductRepo.GetAllProducts - r.Pool.Query: %v", err)
	}
	defer rows.Close()

	var products []entity.Product
	for rows.Next() {
		var product entity.Product
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Quantity,
		)
		if err != nil {
			return nil, fmt.Errorf("ProductRepo.GetAllProducts - rows.Next: %v", err)
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *ProductRepo) GetProductById(ctx context.Context, id int) (entity.Product, error) {
	sql, args, err := squirrel.
		Select("*").
		From("products").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return entity.Product{}, fmt.Errorf("ProductRepo.GetProductById - squirrel.Select: %v", err)
	}

	var product entity.Product
	err = r.Pool.QueryRow(ctx, sql, args).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Quantity,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Product{}, repoerrs.ErrNotFound
		}
		return entity.Product{}, fmt.Errorf("ProductRepo.GetProductById - r.Pool.QueryRow: %v", err)
	}

	return product, nil
}

func (r *ProductRepo) UpdateProduct(ctx context.Context, product entity.Product) error {
	sql, args, err := squirrel.
		Update("products").
		Set("name", product.Name).
		Set("description", product.Description).
		Set("price", product.Price).
		Set("quantity", product.Quantity).
		Where("id = ?", product.ID).
		ToSql()
	if err != nil {
		return fmt.Errorf("ProductRepo.UpdateProduct - squirrel.Update: %v", err)
	}

	if _, err = r.Pool.Exec(ctx, sql, args); err != nil {
		return fmt.Errorf("ProductRepo.UpdateProduct - r.Pool.Exec: %v", err)
	}

	return nil
}

func (r *ProductRepo) DeleteProduct(ctx context.Context, id int) error {
	sql, args, err := squirrel.
		Delete("products").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return fmt.Errorf("ProductRepo.DeleteProduct - squirrel.Delete: %v", err)
	}

	if _, err = r.Pool.Exec(ctx, sql, args); err != nil {
		return fmt.Errorf("ProductRepo.DeleteProduct - r.Pool.Exec: %v", err)
	}

	return nil
}