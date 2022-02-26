package app

import (
	"net/http"

	"github.com/Dryluigi/auction-system/middleware"
	"github.com/gorilla/mux"
)

func setupHandlers(router *mux.Router) {

	router.Use(middleware.EnableCors)
	router.Use(middleware.JsonResponse)

	v1 := router.PathPrefix("/api/v1").Subrouter()

	productsRouter := v1.PathPrefix("/products").Subrouter()
	productsRouter.HandleFunc("", productController.Save).Methods(http.MethodPost)
	productsRouter.HandleFunc("/{id}", productController.Update).Methods(http.MethodPut)
	productsRouter.HandleFunc("/{id}", productController.Delete).Methods(http.MethodDelete)

	bidSessionRouter := v1.PathPrefix("/bid-sessions").Subrouter()
	bidSessionRouter.HandleFunc("", bidSessionController.Save).Methods(http.MethodPost)
	bidSessionRouter.HandleFunc("/{id}", bidSessionController.Delete).Methods(http.MethodDelete)
	bidSessionRouter.HandleFunc("/occuring", bidSessionController.Occuring).Methods(http.MethodGet)
	bidSessionRouter.HandleFunc("/{bidSessionId}/bid/{productId}", bidSessionController.BidItem).Methods(http.MethodPost)
}
