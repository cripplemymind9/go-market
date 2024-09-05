package repository

import (
	"context"

	"github.com/cripplemymind9/go-market/internal/entity"
	"github.com/cripplemymind9/go-market/internal/repository/pgdb"
	"github.com/cripplemymind9/go-market/pkg/postgres"
)

type User interface {
	RegisterUser(ctx context.Context, user entity.User) (int, error)
	LoginUser(ctx context.Context, username string) (entity.User, error)
	GetUserProfile(ctx context.Context, userId int) (entity.User, error)
}

type Product interface {
	AddProduct(ctx context.Context, product entity.Product) (int, error)
	GetAllProducts(ctx context.Context) ([]entity.Product, error)
	GetProductById(ctx context.Context, productId int) (entity.Product, error)
	UpdateProduct(ctx context.Context, product entity.Product) error
	DeleteProduct(ctx context.Context, productId int) error
}

type Purchase interface {
	MakePurchase(ctx context.Context, purchase entity.Purchase) (int, error)
	GetUserPurchases(ctx context.Context, userId int) ([]entity.Purchase, error)
	GetProductPurchases(ctx context.Context, productId int) ([]entity.Purchase, error)
}

type Repositories struct {
	User
	Product
	Purchase
}

func NewRepositories(pg *postgres.Postgres) *Repositories {
	return &Repositories{
		User:     pgdb.NewUserRepo(pg),
		Product:  pgdb.NewProductRepo(pg),
		Purchase: pgdb.NewPurchaseRepo(pg),
	}
}
