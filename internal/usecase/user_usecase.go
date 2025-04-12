package usecase

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"task-manager/internal/model"
)

type UserRepository interface {
	GetAllUsers() ([]model.User, error)
	GetUserByID(id int) (*model.User, error)
	RegisterUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id int) error
}

type UserUsecase interface {
	GetAllUsers() ([]model.User, error)
	GetUserByID(id int) (*model.User, error)
	RegisterUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id int) error
}

type userUsecase struct {
	repo UserRepository
}

func NewUserUsecase(repo UserRepository) UserUsecase {
	return &userUsecase{repo}
}

func (u *userUsecase) GetAllUsers() ([]model.User, error) {
	return u.repo.GetAllUsers()
}

func (u *userUsecase) GetUserByID(id int) (*model.User, error) {
	return u.repo.GetUserByID(id)
}

func (u *userUsecase) RegisterUser(user *model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	user.Password = string(hashedPassword)

	return u.repo.RegisterUser(user)
}

func (u *userUsecase) UpdateUser(user *model.User) error {
	return u.repo.UpdateUser(user)
}

func (u *userUsecase) DeleteUser(id int) error {
	return u.repo.DeleteUser(id)
}
