package database

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Dryluigi/auction-system/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ConnectionAttributes struct {
	DbName     string
	DbHost     string
	DbUser     string
	DbPassword string
	DbPort     uint32
}

func Connect() {
	var err error

	attrs := ConnectionAttributes{}
	err = buildAttributes(&attrs)

	if err != nil {
		log.Fatal("error: " + err.Error())
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		attrs.DbUser,
		attrs.DbPassword,
		attrs.DbHost,
		attrs.DbPort,
		attrs.DbName,
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("database connection can't be established.")
		return
	}

	migrate(DB)

	log.Println("database connection established")
}

func buildAttributes(attrs *ConnectionAttributes) error {
	attrs.DbHost = os.Getenv("DB_HOST")
	attrs.DbUser = os.Getenv("DB_USER")
	attrs.DbPassword = os.Getenv("DB_PASSWORD")
	attrs.DbName = os.Getenv("DB_NAME")

	portStr := os.Getenv("DB_PORT")
	port, err := strconv.Atoi(portStr)

	if err != nil {
		return errors.New("provide valid port number")
	}

	attrs.DbPort = uint32(port)

	return nil
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.BidProduct{},
		&model.BidSession{},
		&model.Product{},
	)
}
