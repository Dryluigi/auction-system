package v1

type BidSessionSave struct {
	TimeStart string `json:"time_start" validate:"required,datetime"`
	TimeEnd   string `json:"time_end" validate:"required,datetime"`
}
