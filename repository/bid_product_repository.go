package repository

import "github.com/Dryluigi/auction-system/model"

type BidProductRepository interface {
	BidItem(bidSessionId uint, productId uint) (*model.BidProduct, error)
}
