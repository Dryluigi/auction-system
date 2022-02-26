package repository

import (
	"time"

	"github.com/Dryluigi/auction-system/model"
)

type BidSessionRepository interface {
	Save(model *model.BidSession) error
	Update(model *model.BidSession) error
	DeleteById(id uint) error
	GetOccuring(model *model.BidSession, now time.Time) error
	VerifyExist(id uint) bool
}
