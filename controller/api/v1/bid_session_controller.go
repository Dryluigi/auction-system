package v1

import "net/http"

type BidSessionController interface {
	Save(w http.ResponseWriter, r *http.Request)
	Occuring(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	BidItem(w http.ResponseWriter, r *http.Request)
}
