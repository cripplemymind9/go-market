package impl

import (
	"github.com/cripplemymind9/go-market/internal/repository/repoerrs"
	"github.com/cripplemymind9/go-market/internal/service/serviceerrs"
	"github.com/cripplemymind9/go-market/internal/service/types"
	"github.com/cripplemymind9/go-market/internal/repository"
	"github.com/cripplemymind9/go-market/internal/entity"
	log "github.com/sirupsen/logrus"
	"context"
	"errors"
)

type ProductService struct {
	productRepo repository.Product
}

func NewProductService(productRepo repository.Product) *ProductService {
	return &ProductService{productRepo: productRepo}
}

func (s *ProductService) AddProduct(ctx context.Context, input types.ProductAddProductInput) (int, error) {
	product := entity.Product{
		Name: input.Name,
		Description: input.Description,
		Price: input.Price,
		Quantity: input.Quantity,
	}
	
	id, err := s.productRepo.AddProduct(ctx, product)
	if err != nil {
		if errors.Is(err, repoerrs.ErrAlreadyExists) {
			return 0, serviceerrs.ErrProductAlreadyExists
		}
		log.Errorf("ProductService.AddProduct - s.productRepo.AddProduct: %v", err)
		return 0, serviceerrs.ErrCannotCreateProduct
	}

	return id, nil
}

func (s *ProductService) GetAllProducts(ctx context.Context) ([]entity.Product, error) {
	products, err := s.productRepo.GetAllProducts(ctx)
	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return nil, serviceerrs.ErrNoProductsAvailable
		}
		log.Errorf("ProductService.GetAllProducts - s.productRepo.GetAllProducts: %v", err)
		return nil, serviceerrs.ErrCannotGetProducts
	}

	return products, nil
}

func (s *ProductService) GetProductById(ctx context.Context, productId int) (entity.Product, error) {
	return s.productRepo.GetProductById(ctx, productId)
}

func (s *ProductService) UpdateProduct(ctx context.Context, input types.ProductUpdateProductInput) error {
	product := entity.Product{
		ID: input.ID,
		Name: input.Name,
		Description: input.Description,
		Price: input.Price,
		Quantity: input.Quantity,
	}

	return s.productRepo.UpdateProduct(ctx, product)
}

func (s *ProductService) DeleteProduct(ctx context.Context, productId int) error {
	return s.productRepo.DeleteProduct(ctx, productId)
}