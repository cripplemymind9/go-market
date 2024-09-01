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

func (r *PurchaseRepo) MakePurchase(ctx context.Context, purchase entity.Purchase) (int, error) {
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("PurchaseRepo.MakePurchase - r.Pool.Begin: %v", err)
	}
	defer tx.Rollback(ctx)
	
	sql, args, err := r.Builder.
		Update("products").
		Set("quantity", squirrel.Expr("quantity - ?", purchase.Quantity)).
		Where("id = ?", purchase.ProductID).
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("PurchaseRepo.MakePurchase - r.Builder..Update: %v", err)
	}

	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		return 0, fmt.Errorf("PurchaseRepo.MakePurchase - tx.Exec: %v", err)
	}

	sql, args, err = r.Builder.
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
		return 0, fmt.Errorf("PurchaseRepo.MakePurchase - r.Builder.Insert:%v", err)
	}

	var id int
	err = tx.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok {
			if pgErr.Code == "23505" {
				return 0, repoerrs.ErrAlreadyExists
			}
		}
		return 0, fmt.Errorf("PurchaseRepo.MakePurchase - tx.QueryRow:%v", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return 0, fmt.Errorf("PurchaseRepo.MakePurchase - tx.Commit:%v", err)
	}

	return id, nil
}

func (r *PurchaseRepo) GetUserPurchases(ctx context.Context, userId int) ([]entity.Purchase, error) {
	sql, args, err := r.Builder.
		Select("*").
		From("purchases").
		Where("user_id = ?", userId).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("PurchaseRepo.GetUserPurchases - r.Builder.Select:%v", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("PurchaseRepo.GetUserPurchases - r.Pool.Query:%v", err)
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
			return nil, fmt.Errorf("PurchaseRepo.GetUserPurchases - rows.Next:%v", err)
		}
		purchases = append(purchases, purchase)
	}

	return purchases, nil
}

func (r *PurchaseRepo) GetProductPurchases(ctx context.Context, productId int) ([]entity.Purchase, error) {
	sql, args, err := r.Builder.
		Select("*").
		From("purchases").
		Where("product_id = ?", productId).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("PurchaseRepo.GetProductPurchases - r.Builder.Select:%v", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("PurchaseRepo.GetProductPurchases - r.Pool.Query:%v", err)
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
			return nil, fmt.Errorf("PurchaseRepo.GetProductPurchases - rows.Next:%v", err)
		}
		purchases = append(purchases, purchase)
	}	

	return purchases, nil
}