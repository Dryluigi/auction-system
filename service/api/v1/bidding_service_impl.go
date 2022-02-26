package v1

import (
	"errors"

	"github.com/Dryluigi/auction-system/model"
	"github.com/Dryluigi/auction-system/repository"
	serviceError "github.com/Dryluigi/auction-system/service/api/errors"
	"github.com/Dryluigi/auction-system/service/firebase"
	"gorm.io/gorm"
)

type BiddingServiceImpl struct {
	BidProductRepository repository.BidProductRepository
}

func (service *BiddingServiceImpl) BidItem(bidSessionId uint, productId uint) (*model.BidProduct, error) {
	var realtimeDb firebase.RealtimeDatabase = &firebase.RealtimeDatabaseImpl{}
	bidProduct, err := service.BidProductRepository.BidItem(bidSessionId, productId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, serviceError.ErrEntityNotFound
		} else {
			return nil, err
		}
	}

	realtimeDb.Save("bid_products", bidProduct)

	return bidProduct, nil
}
