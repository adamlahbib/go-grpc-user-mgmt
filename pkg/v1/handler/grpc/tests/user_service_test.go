package grpc_test

import (
	"context"
	"testing"

	pb "github.com/adamlahbib/go-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestCreateUser(t *testing.T) {
	// establish a connection to the server
	conn, err := grpc.NewClient(
		"localhost:5001",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatalf("Failed to dial the server: %v", err)
	}
	defer conn.Close()

	// create a new client
	client := pb.NewUserServiceClient(conn)

	// send a request to the server
	req := &pb.CreateUserRequest{
		Name:  "Adam Lahbib",
		Email: "contactme@adamlahbib.me",
	}

	res, err := client.Create(context.Background(), req)
	if err != nil {
		t.Fatalf("Failed to create a user: %v", err)
	}

	// verify the response
	if res.GetName() != req.GetName() {
		t.Fatalf("Expected name %s, got %s", req.GetName(), res.GetName())
	}

	if res.GetEmail() != req.GetEmail() {
		t.Fatalf("Expected email %s, got %s", req.GetEmail(), res.GetEmail())
	}

	if res.GetId() == "" {
		t.Fatalf("Expected id to be not empty")
	}
}
