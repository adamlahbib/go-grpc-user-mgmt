package usecase

import (
	"errors"
	"fmt"

	"github.com/adamlahbib/go-grpc/internal/models"
	interfaces "github.com/adamlahbib/go-grpc/pkg/v1"
	"gorm.io/gorm"
)

type UseCase struct {
	repo interfaces.RepoInterface
}

func New(repo interfaces.RepoInterface) interfaces.UseCaseInterface {
	return &UseCase{repo}
}

// Create a new user with supplied data
func (u *UseCase) Create(user models.User) (models.User, error) {
	// check if email already exists
	if err := u.repo.GetByEmail(user.Email); !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.User{}, errors.New("email already exists")
	}

	// proceeding to create the user in case email does not exist
	return u.repo.Create(user)
}

// Retrieve the user instance by id
func (u *UseCase) Get(id string) (models.User, error) {
	var user models.User
	var err error

	// check if user exists
	if user, err = u.repo.Get(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, errors.New("user not found")
		}

		// return the error if it is not a record not found error
		return models.User{}, err
	}

	// return the user instance otherwise
	return user, nil
}

// Update the user instance
func (u *UseCase) Update(modifiedUser models.User) error {
	var user models.User
	var err error

	// check if user exists
	if user, err = u.repo.Get(fmt.Sprint(modifiedUser.ID)); err != nil {
		return err
	}

	// check if the email is not being updated, only name can be updated
	if user.Email != modifiedUser.Email {
		return errors.New("email cannot be changed")
	}

	// proceed with the update
	return u.repo.Update(modifiedUser)
}

// Delete the user whose ID is supplied
func (u *UseCase) Delete(id string) error {
	// check if user exists
	if _, err := u.repo.Get(id); err != nil {
		return err
	}

	return u.repo.Delete(id)
}
