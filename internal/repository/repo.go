package repository

import (
	"github.com/cripplemymind9/go-market/internal/repository/pgdb"
	"github.com/cripplemymind9/go-market/internal/entity"
	"github.com/cripplemymind9/go-market/pkg/postgres"
	"context"
)

type User interface {
	CreateUser(ctx context.Context, user entity.User) (int, error)
	GetUserByUserId(ctx context.Context, id int) (entity.User, error)
	GetUserByUsername(ctx context.Context, username string) (entity.User, error)
	GetUserByUsernameAndPassword(ctx context.Context, username, password string) (entity.User, error)
}

type Product interface {
	CreateProduct(ctx context.Context, product entity.Product) (int, error)
	GetAllProducts(ctx context.Context) ([]entity.Product, error)
	GetProductById(ctx context.Context, id int) (entity.Product, error)
	UpdateProduct(ctx context.Context, product entity.Product) error
	DeleteProduct(ctx context.Context, id int) error
}

type Purchase interface {
	CreatePurchase(ctx context.Context, purchase entity.Purchase) (int, error)
	GetPurchasesByUserId(ctx context.Context, id int) ([]entity.Purchase, error)
	GetPurchasesByProductId(ctx context.Context, id int) ([]entity.Purchase, error)
}

type Repositories struct {
	User
	Product
	Purchase
}

func NewRepositories(pg *postgres.Postgres) *Repositories {
	return &Repositories{
		User: pgdb.NewUserRepo(pg),
		Product: pgdb.NewProductRepo(pg),
	}
}