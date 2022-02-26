package v1

import (
	"github.com/Dryluigi/auction-system/model"
	request "github.com/Dryluigi/auction-system/request/v1"
)

type BidSessionService interface {
	Save(request *request.BidSessionSave) (*model.BidSession, error)
	GetOccuring() (*model.BidSession, error)
	Delete(id uint) error
}
