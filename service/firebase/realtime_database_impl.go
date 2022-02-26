package firebase

import (
	"context"
	"log"
	"os"

	firebasePckg "firebase.google.com/go"
	"google.golang.org/api/option"
)

type RealtimeDatabaseImpl struct {
}

func (db *RealtimeDatabaseImpl) Save(key string, data interface{}) error {

	ctx := context.Background()

	cfg := &firebasePckg.Config{
		DatabaseURL: os.Getenv("FIREBASE_REALTIME_DATABASE_URL"),
	}
	opt := option.WithCredentialsFile("../../" + os.Getenv("FIREBASE_SERVICE_ACCOUNT_PATH"))

	app, err := firebasePckg.NewApp(ctx, cfg, opt)
	if err != nil {
		log.Fatalln("Error initializing app:", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}

	ref := client.NewRef(key)

	if _, err := ref.Push(ctx, data); err != nil {
		log.Fatalln("Error pushing child node:", err)
	}

	return nil
}

func (db *RealtimeDatabaseImpl) Update(key string, data interface{}) error {
	return nil
}
