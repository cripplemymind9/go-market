package impl

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"

	"github.com/cripplemymind9/go-market/internal/repository/repoerrs"
	"github.com/cripplemymind9/go-market/internal/service/serviceerrs"
	"github.com/cripplemymind9/go-market/internal/service/types"
	"github.com/cripplemymind9/go-market/internal/repository"
	"github.com/cripplemymind9/go-market/internal/entity"
)

type PurchaseService struct {
	purchaseRepo repository.Purchase
}

func NewPurchaseService(purchaseRepo repository.Purchase) *PurchaseService {
	return &PurchaseService{purchaseRepo: purchaseRepo}
}

func (s *PurchaseService) MakePurchase(ctx context.Context, input types.PurchaseMakePurchaseInput) (int, error) {
	purchase := entity.Purchase{
		UserID: input.UserID,
		ProductID: input.ProductID,
		Quantity: input.Quantity,
	}

	id, err := s.purchaseRepo.MakePurchase(ctx, purchase)
	if err != nil {
		log.Errorf("PurchaseService.MakePurchase - s.purchaseRepo.MakePurchase: %v", err)
		return 0, serviceerrs.ErrCannotCreatePurchase
	}

	return id, nil
}

func (s *PurchaseService) GetUserPurchases(ctx context.Context, userId int) ([]entity.Purchase, error) {
	purchases, err := s.purchaseRepo.GetUserPurchases(ctx, userId)
	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return nil, serviceerrs.ErrNoUserPurchasesFound
		}
		log.Errorf("PurchaseService.GetUserPurchases - s.purchaseRepo.GetUserPurchases: %v", err)
		return nil, serviceerrs.ErrCannotGetUserPurchases
	}

	return purchases, nil
}

func (s *PurchaseService) GetProductPurchases(ctx context.Context, productId int) ([]entity.Purchase, error) {
	purchases, err := s.purchaseRepo.GetProductPurchases(ctx, productId)
	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return nil, serviceerrs.ErrNoProductPurchasesFound
		}
		log.Errorf("PurchaseService.GetProductPurchases - s.purchaseRepo.GetProductPurchases: %v", err)
		return nil, serviceerrs.ErrCannotGetProductPurchases
	}

	return purchases, nil
}