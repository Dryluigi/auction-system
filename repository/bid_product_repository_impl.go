package repository

import (
	"github.com/Dryluigi/auction-system/model"
	"gorm.io/gorm"
)

type BidProductRepositoryImpl struct {
	DB *gorm.DB
}

func (repository *BidProductRepositoryImpl) BidItem(bidSessionId uint, productId uint) (*model.BidProduct, error) {
	bidProduct := &model.BidProduct{}
	err := repository.DB.Where("bid_session_id = ? AND product_id = ?", bidSessionId, productId).Preload("Product").First(bidProduct).Error

	if err != nil {
		return nil, err
	}

	bidProduct.StartPrice += bidProduct.BidIncrement

	if err = repository.DB.Save(bidProduct).Error; err != nil {
		return nil, err
	}

	return bidProduct, nil
}
