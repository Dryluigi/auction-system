package repository

import (
	"time"

	"github.com/Dryluigi/auction-system/model"
	"gorm.io/gorm"
)

type BidSessionRepositoryImpl struct {
	DB *gorm.DB
}

func (bs *BidSessionRepositoryImpl) Save(model *model.BidSession) error {
	if err := bs.DB.Create(model).Error; err != nil {
		return err
	}

	return nil
}

func (bs *BidSessionRepositoryImpl) Update(model *model.BidSession) error {
	if err := bs.DB.Save(model).Error; err != nil {
		return err
	}

	return nil
}

func (bs *BidSessionRepositoryImpl) DeleteById(id uint) error {
	if err := bs.DB.Delete(&model.BidSession{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (bs *BidSessionRepositoryImpl) GetOccuring(model *model.BidSession, now time.Time) error {
	nowStr := now.Format(time.RFC3339)
	if err := bs.DB.Where("time_start <= ? AND time_end > ?", nowStr, nowStr).Preload("BidProducts").Preload("BidProducts.Product").First(model).Error; err != nil {
		return err
	}

	return nil
}

func (bs *BidSessionRepositoryImpl) VerifyExist(id uint) bool {
	model := &model.BidSession{}
	err := bs.DB.First(model, id).Error

	return err == nil
}
