package model

import "time"

type BidSession struct {
	ID          uint          `json:"id" gorm:"autoIncrement"`
	TimeStart   time.Time     `json:"time_start" gorm:"not null"`
	TimeEnd     time.Time     `json:"time_end" gorm:"not null"`
	CreatedAt   time.Time     `json:"created_at" gorm:"<-:create"`
	UpdatedAt   time.Time     `json:"updated_at"`
	BidProducts []*BidProduct `json:"products,omitempty"`
}

func (bs *BidSession) TableName() string {
	return "tbl_bid_sessions"
}
