package v1

import (
	"github.com/Dryluigi/auction-system/model"
	"github.com/Dryluigi/auction-system/repository"
	request "github.com/Dryluigi/auction-system/request/v1"
	serviceErrors "github.com/Dryluigi/auction-system/service/api/errors"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
}

func (service *ProductServiceImpl) Save(p *request.ProductSave) (*model.Product, error) {
	modelProduct := &model.Product{
		Name:  p.Name,
		Price: p.Price,
	}
	err := service.ProductRepository.Save(modelProduct)

	if err != nil {
		return nil, err
	}

	return modelProduct, err
}

func (service *ProductServiceImpl) Update(id uint, p *request.ProductUpdate) (*model.Product, error) {
	if !service.ProductRepository.ProductExists(id) {
		return nil, serviceErrors.ErrEntityNotFound
	}

	modelProduct := &model.Product{
		ID:    uint(id),
		Name:  p.Name,
		Price: p.Price,
	}

	err := service.ProductRepository.Update(modelProduct)

	if err != nil {
		return nil, err
	}

	return modelProduct, nil
}

func (service *ProductServiceImpl) Delete(id uint) error {
	if !service.ProductRepository.ProductExists(id) {
		return serviceErrors.ErrEntityNotFound
	}

	err := service.ProductRepository.DeleteById(id)

	if err != nil {
		return err
	}

	return nil
}
