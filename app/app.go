package app

import (
	"github.com/Dryluigi/auction-system/database"
	"github.com/Dryluigi/auction-system/validator"
	"github.com/gorilla/mux"
)

func New() *mux.Router {
	app := mux.NewRouter()

	database.Connect()
	validator.InitValidator()
	doInjection()
	setupHandlers(app)

	return app
}
