package main

import (
	"log"
	"net/http"

	"github.com/Dryluigi/auction-system/app"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	app := app.New()

	log.Fatal(http.ListenAndServe(":8080", app))
}
