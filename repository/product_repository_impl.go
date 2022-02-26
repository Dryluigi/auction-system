package repository

import (
	"github.com/Dryluigi/auction-system/model"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func (p *ProductRepositoryImpl) Save(product *model.Product) error {
	if err := p.DB.Create(&product).Error; err != nil {
		return err
	}

	return nil
}

func (p *ProductRepositoryImpl) FindById(id uint) (*model.Product, error) {
	product := &model.Product{}
	if err := p.DB.First(product, id).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductRepositoryImpl) Update(product *model.Product) error {
	if err := p.DB.Save(product).Error; err != nil {
		return err
	}

	return nil
}

func (p *ProductRepositoryImpl) DeleteById(id uint) error {
	if err := p.DB.Delete(&model.Product{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (p *ProductRepositoryImpl) ProductExists(id uint) bool {
	count := int64(0)
	err := p.DB.Model(&model.Product{}).
		Where("id = ?", id).
		Count(&count).
		Error

	if err != nil {
		return false
	}

	return count > 0
}
