package model

import "time"

type Product struct {
	ID        uint      `json:"id" gorm:"autoIncrement"`
	Name      string    `json:"name" gorm:"not null"`
	Price     int       `json:"price" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"<-:create"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *Product) TableName() string {
	return "tbl_products"
}
