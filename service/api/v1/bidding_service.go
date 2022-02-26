package v1

import "github.com/Dryluigi/auction-system/model"

type BiddingService interface {
	BidItem(bidSessionId uint, productId uint) (*model.BidProduct, error)
}
