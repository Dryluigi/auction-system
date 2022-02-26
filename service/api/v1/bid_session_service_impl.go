package v1

import (
	"errors"
	"time"

	"github.com/Dryluigi/auction-system/model"
	"github.com/Dryluigi/auction-system/repository"
	request "github.com/Dryluigi/auction-system/request/v1"
	serviceError "github.com/Dryluigi/auction-system/service/api/errors"
	"gorm.io/gorm"
)

type BidSessionServiceImpl struct {
	BidSessionRepository repository.BidSessionRepository
}

func (service *BidSessionServiceImpl) Save(request *request.BidSessionSave) (*model.BidSession, error) {
	timeStart, err := time.Parse(time.RFC3339, request.TimeStart)
	if err != nil {
		return nil, err
	}

	timeEnd, err := time.Parse(time.RFC3339, request.TimeEnd)
	if err != nil {
		return nil, err
	}

	model := &model.BidSession{
		TimeStart: timeStart,
		TimeEnd:   timeEnd,
	}

	err = service.BidSessionRepository.Save(model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (service *BidSessionServiceImpl) GetOccuring() (*model.BidSession, error) {
	now := time.Now()

	model := &model.BidSession{}
	if err := service.BidSessionRepository.GetOccuring(model, now); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return model, nil
}

func (service *BidSessionServiceImpl) Delete(id uint) error {
	if !service.BidSessionRepository.VerifyExist(id) {
		return serviceError.ErrEntityNotFound
	}
	err := service.BidSessionRepository.DeleteById(id)

	if err != nil {
		return err
	}

	return nil
}
