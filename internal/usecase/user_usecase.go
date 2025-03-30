package usecase

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"task-manager/internal/models"
)

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id int) (*models.User, error)
	RegisterUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
}

type UserUsecase interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id int) (*models.User, error)
	RegisterUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
}

type userUsecase struct {
	repo UserRepository
}

func NewUserUsecase(repo UserRepository) UserUsecase {
	return &userUsecase{repo}
}

func (u *userUsecase) GetAllUsers() ([]models.User, error) {
	return u.repo.GetAllUsers()
}

func (u *userUsecase) GetUserByID(id int) (*models.User, error) {
	return u.repo.GetUserByID(id)
}

func (u *userUsecase) RegisterUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	user.Password = string(hashedPassword)

	return u.repo.RegisterUser(user)
}

func (u *userUsecase) UpdateUser(user *models.User) error {
	return u.repo.UpdateUser(user)
}

func (u *userUsecase) DeleteUser(id int) error {
	return u.repo.DeleteUser(id)
}
