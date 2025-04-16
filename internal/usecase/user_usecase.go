package usecase

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"task-manager/internal/model"
)

type UserRepository interface {
	GetAllUsers() ([]model.User, error)
	GetUserByID(id uint) (*model.User, error)
	RegisterUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id uint) error
	GetUserByEmail(email string) (*model.User, error)
	CheckPassword(email, password string) (*model.User, error)
}

type userUsecase struct {
	repo UserRepository
}

func NewUserUsecase(repo UserRepository) *userUsecase {
	return &userUsecase{repo}
}

func (u *userUsecase) GetAllUsers() ([]model.User, error) {
	return u.repo.GetAllUsers()
}

func (u *userUsecase) GetUserByID(id uint) (*model.User, error) {
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

func (u *userUsecase) DeleteUser(id uint) error {
	return u.repo.DeleteUser(id)
}

func (u *userUsecase) GetUserByEmail(email string) (*model.User, error) {
	return u.repo.GetUserByEmail(email)
}

func (u *userUsecase) CheckPassword(email, password string) (*model.User, error) {
	user, err := u.repo.GetUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, fmt.Errorf("invalid credentials")
		}
		return nil, fmt.Errorf("failed to check password: %w", err)
	}

	user.Password = ""
	return user, nil
}
