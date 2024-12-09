package main

import (
	"log"
	"net"

	dbConfig "github.com/adamlahbib/go-grpc/internal/db"
	"github.com/adamlahbib/go-grpc/internal/models"
	handler "github.com/adamlahbib/go-grpc/pkg/v1/handler/grpc"
	repo "github.com/adamlahbib/go-grpc/pkg/v1/repository"
	useCase "github.com/adamlahbib/go-grpc/pkg/v1/usecase"
	"google.golang.org/grpc"
)

func main() {
	// Connect to the database
	db := dbConfig.DbConn()

	// Migrate the database
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("failed to migrate the database: %v", err)
	}

	// create a listener for the server
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("Error creating the listener: %v", err)
	}

	// create a new gRPC server instance
	server := grpc.NewServer()
	handler.NewServer(server, useCase.New(repo.New(db)))

	// start serving
	log.Fatal(server.Serve(lis))
}
