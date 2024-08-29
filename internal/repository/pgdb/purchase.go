package pgdb

import (
	"github.com/cripplemymind9/go-market/internal/repository/repoerrs"
	"github.com/cripplemymind9/go-market/internal/entity"
	"github.com/cripplemymind9/go-market/pkg/postgres"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgconn"
	"context"
	"errors"
	"fmt"
)

type PurchaseRepo struct {
	*postgres.Postgres
}

func NewPurchaseRepo(pg *postgres.Postgres) *PurchaseRepo {
	return &PurchaseRepo{pg}
}

func (r *PurchaseRepo) CreatePurchase(ctx context.Context, purchase entity.Purchase) (int, error) {
	sql, args, err := squirrel.
		Insert("purchases").
		Columns("user_id", "product_id", "quantity").
		Values(
			purchase.UserID,
			purchase.ProductID,
			purchase.Quantity,
		).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("PurchaseRepo.CreatePurchase - squirrel.Insert:%v", err)
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
		return 0, fmt.Errorf("PurchaseRepo.CreatePurchase - r.Pool.QueryRow:%v", err)
	}

	return id, nil
}

func (r *PurchaseRepo) GetPurchasesByUserId(ctx context.Context, id int) ([]entity.Purchase, error) {
	sql, args, err := squirrel.
		Select("*").
		From("purchases").
		Where("user_id = ?", id).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("PurchaseRepo.GetPurchasesByUserId - squirrel.Select:%v", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("PurchaseRepo.GetPurchasesByUserId - r.Pool.Query:%v", err)
	}
	defer rows.Close()

	var purchases []entity.Purchase
	for rows.Next() {
		var purchase entity.Purchase
		err = rows.Scan(
			&purchase.ID,
			&purchase.UserID,
			&purchase.ProductID,
			&purchase.Quantity,
			&purchase.Timestamp,
		)
		if err != nil {
			return nil, fmt.Errorf("PurchaseRepo.GetPurchasesByUserId - rows.Next:%v", err)
		}
		purchases = append(purchases, purchase)
	}

	return purchases, nil
}

func (r *PurchaseRepo) GetPurchasesByProductId(ctx context.Context, id int) ([]entity.Purchase, error) {
	sql, args, err := squirrel.
		Select("*").
		From("purchases").
		Where("product_id = ?", id).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("PurchaseRepo.GetPurchasesByProductId - squirrel.Select:%v", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("PurchaseRepo.GetPurchasesByProductId - r.Pool.Query:%v", err)
	}
	defer rows.Close()

	var purchases []entity.Purchase
	for rows.Next() {
		var purchase entity.Purchase
		err = rows.Scan(
			&purchase.ID,
			&purchase.UserID,
			&purchase.ProductID,
			&purchase.Quantity,
			&purchase.Timestamp,
		)
		if err != nil {
			return nil, fmt.Errorf("PurchaseRepo.GetPurchasesByProductId - rows.Next:%v", err)
		}
		purchases = append(purchases, purchase)
	}	

	return purchases, nil
}