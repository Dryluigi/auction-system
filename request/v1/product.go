package v1

type ProductSave struct {
	Name  string `json:"name" validate:"required"`
	Price int    `json:"price" validate:"required"`
}

type ProductUpdate struct {
	Name  string `json:"name" validate:"required"`
	Price int    `json:"price" validate:"required"`
}
