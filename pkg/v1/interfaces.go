package v1

import "github.com/adamlahbib/go-grpc/internal/models"

// an interface for repo methods
type RepoInterface interface {
	// create a new user with supplied data
	Create(models.User) (models.User, error)

	// retrieve the user instance by id
	Get(id string) (models.User, error)

	// update the user instance
	Update(models.User) error

	// delete the user whose ID is supplied
	Delete(id string) error

	// fetch user by email, to test the unique constraint in usecase
	GetByEmail(email string) error
}

// an interface for service methods
type UseCaseInterface interface {
	Create(models.User) (models.User, error)
	Get(id string) (models.User, error)
	Update(models.User) error
	Delete(id string) error
}
