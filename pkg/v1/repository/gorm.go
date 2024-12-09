package repo

import (
	"github.com/adamlahbib/go-grpc/internal/models"
	interfaces "github.com/adamlahbib/go-grpc/pkg/v1"
	"gorm.io/gorm"
)

// Repo is a struct that defines the repository
type Repo struct {
	db *gorm.DB
}

// New creates a new instance of the Repo
func New(db *gorm.DB) interfaces.RepoInterface {
	return &Repo{db}
}

// Create a new user with supplied data
func (r *Repo) Create(user models.User) (models.User, error) {
	return user, r.db.Create(&user).Error
}

// Retrieve the user instance by id
func (r *Repo) Get(id string) (models.User, error) {
	var user models.User
	return user, r.db.First(&user, id).Error
}

// Update the user instance
func (r *Repo) Update(user models.User) error {
	var u models.User
	if err := r.db.First(&u, user.ID).Error; err != nil {
		return err
	}
	u.Name = user.Name
	return r.db.Save(&u).Error
}

// Delete the user whose ID is supplied
func (r *Repo) Delete(id string) error {
	return r.db.Delete(&models.User{}, id).Error
}

// Fetch user by email, to test the unique constraint in usecase
func (r *Repo) GetByEmail(email string) error {
	var user models.User
	return r.db.First(&user, "email = ?", email).Error
}
