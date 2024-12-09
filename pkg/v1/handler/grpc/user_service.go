package grpc

import (
	"context"
	"errors"
	"fmt"

	"github.com/adamlahbib/go-grpc/internal/models"
	interfaces "github.com/adamlahbib/go-grpc/pkg/v1"

	pb "github.com/adamlahbib/go-grpc/proto"
	"google.golang.org/grpc"
)

type UserService struct {
	useCase interfaces.UseCaseInterface
	pb.UnimplementedUserServiceServer
}

func NewServer(grpcServer *grpc.Server, usecase interfaces.UseCaseInterface) {
	userGrpc := &UserService{useCase: usecase}
	pb.RegisterUserServiceServer(grpcServer, userGrpc)
}

// Create a new user function through the CreateUserRquest message of Proto
func (s *UserService) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserProfileResponse, error) {
	// convert the request to a user model
	data := models.User{
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}

	// verify if the supplied data is not empty
	if data.Name == "" || data.Email == "" {
		return &pb.UserProfileResponse{}, errors.New("name and email are required")
	}

	// send the data to the usecase
	user, err := s.useCase.Create(data)
	if err != nil {
		return &pb.UserProfileResponse{}, err
	}

	// format the response as a UserProfileResponse message
	return &pb.UserProfileResponse{
		Id:    fmt.Sprint(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
