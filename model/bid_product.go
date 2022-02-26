package model

import "time"

type BidProduct struct {
	ID              uint       `json:"id" gorm:"autoIncrement"`
	StartPrice      int        `json:"start_price"`
	EndPrice        int        `json:"end_price"`
	CurrentBidPrice int        `json:"current_bid_price"`
	BidIncrement    int        `json:"bid_increment"`
	ProductId       uint       `json:"product_id"`
	BidSessionId    uint       `json:"bid_session_id"`
	CreatedAt       time.Time  `json:"created_at" gorm:"<-:create"`
	UpdatedAt       time.Time  `json:"updated_at"`
	Product         Product    `json:"product,omitempty"`
	BidSession      BidSession `json:"-" gorm:"foreignKey:ID;references:BidSessionId"`
}

func (bs *BidProduct) TableName() string {
	return "tbl_bid_products"
}
