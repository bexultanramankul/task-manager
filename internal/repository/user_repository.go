// repository/user_gorm.go
package repository

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"task-manager/internal/model"
	"task-manager/pkg/logger"

	"gorm.io/gorm"
)

type userRepoGorm struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepoGorm {
	return &userRepoGorm{db}
}

func (r *userRepoGorm) GetAllUsers() ([]model.User, error) {
	var users []model.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepoGorm) GetUserByID(id uint) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}

func (r *userRepoGorm) RegisterUser(user *model.User) error {
	var existing model.User
	if err := r.db.Where("email = ?", user.Email).First(&existing).Error; err == nil {
		log.Printf("User registration failed: email %s already exists", user.Email)
		return errors.New("user with this email already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("failed to check existing user: %w", err)
	}

	if err := r.db.Create(user).Error; err != nil {
		log.Printf("Error registering user %s: %v", user.Email, err)
		return fmt.Errorf("failed to register user: %w", err)
	}

	logger.Log.Printf("User %d registered successfully", user.ID)
	return nil
}

func (r *userRepoGorm) UpdateUser(user *model.User) error {
	result := r.db.Model(&model.User{}).
		Where("id = ?", user.ID).
		Updates(map[string]interface{}{
			"username":   user.Username,
			"email":      user.Email,
			"updated_at": gorm.Expr("NOW()"),
		})

	if result.Error != nil {
		return fmt.Errorf("failed to update user: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (r *userRepoGorm) DeleteUser(id uint) error {
	result := r.db.Delete(&model.User{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete user: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}

	logger.Log.Printf("User %d deleted successfully", id)
	return nil
}

func (r *userRepoGorm) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}

func (r *userRepoGorm) CheckPassword(email, password string) (*model.User, error) {
	user, err := r.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	// Здесь должна быть проверка хеша пароля
	// Например, если вы используете bcrypt:
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Временно просто сравниваем строки (небезопасно!)
	if user.Password != password {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
