package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConn() *gorm.DB {
	db, err := gorm.Open(
		postgres.Open(
			"host=localhost port=5432 user=postgres dbname=go_grpc password=postgres sslmode=disable",
		),
		&gorm.Config{},
	)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	return db
}
