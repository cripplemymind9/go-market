package service

import (
	"github.com/cripplemymind9/go-market/internal/entity"
	"github.com/cripplemymind9/go-market/internal/repository"
	"github.com/cripplemymind9/go-market/internal/service/impl"
	"github.com/cripplemymind9/go-market/internal/service/types"
	"github.com/cripplemymind9/go-market/pkg/hasher"
	"context"
	"time"
)

type Auth interface {
	RegisterUser(ctx context.Context, input types.AuthRegisterUserInput) (int, error)
	GenerateToken(ctx context.Context, input types.AuthGenerateTokenInput) (string, error)
	ParseToken(token string) (int, error)
}

type Product interface {
	AddProduct(ctx context.Context, input types.ProductAddProductInput) (int, error)
	GetAllProducts(ctx context.Context) ([]entity.Product, error)
	GetProductById(ctx context.Context, productId int) (entity.Product, error)
	UpdateProduct(ctx context.Context, input types.ProductUpdateProductInput) error
	DeleteProduct(ctx context.Context, productId int) error
}

type Purchase interface {
	MakePurchase(ctx context.Context, input types.PurchaseMakePurchaseInput) (int, error)
	GetUserPurchases(ctx context.Context, userId int) ([]entity.Purchase, error)
	GetProductPurchases(ctx context.Context, productId int) ([]entity.Purchase, error)
}

type Services struct {
	Auth Auth
	Product Product
	Purchase Purchase
}

type ServiceDependencies struct {
	Repos repository.Repositories
	Hasher hasher.PasswordHasher

	SignKey string
	TokenTTL time.Duration
}

func NewServices(deps ServiceDependencies) *Services {
	return &Services{
		Auth: impl.NewAuthService(deps.Repos.User, deps.Hasher, deps.SignKey, deps.TokenTTL),
		Product: impl.NewProductService(deps.Repos.Product),
		Purchase: impl.NewPurchaseService(deps.Repos.Purchase),
	}
}