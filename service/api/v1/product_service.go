package v1

import (
	"github.com/Dryluigi/auction-system/model"
	request "github.com/Dryluigi/auction-system/request/v1"
)

type ProductService interface {
	Save(p *request.ProductSave) (*model.Product, error)
	Update(id uint, p *request.ProductUpdate) (*model.Product, error)
	Delete(id uint) error
}
