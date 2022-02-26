package repository

import "github.com/Dryluigi/auction-system/model"

type ProductRepository interface {
	Save(product *model.Product) error
	FindById(id uint) (*model.Product, error)
	Update(product *model.Product) error
	DeleteById(id uint) error
	ProductExists(id uint) bool
}
