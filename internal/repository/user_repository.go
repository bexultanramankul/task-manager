package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"task-manager/internal/model"
)

type UserRepoImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepoImpl {
	return &UserRepoImpl{db}
}

func (r *UserRepoImpl) GetAllUsers() ([]model.User, error) {
	rows, err := r.db.Query("SELECT id, username, email, created_at, updated_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepoImpl) GetUserByID(id int) (*model.User, error) {
	var user model.User
	err := r.db.QueryRow(
		"SELECT id, username, email, created_at, updated_at FROM users WHERE id = $1", id,
	).Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &user, nil
}

func (r *UserRepoImpl) RegisterUser(user *model.User) error {
	err := r.db.QueryRow(`
		INSERT INTO users (username, email, password_hash, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW())
		ON CONFLICT (email) DO NOTHING
		RETURNING id, created_at, updated_at
	`, user.Username, user.Email, user.Password).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		log.Printf("Error registering user %s: %v", user.Email, err)
		return fmt.Errorf("failed to register user: %w", err)
	}

	if user.ID == 0 {
		log.Printf("User registration failed: email %s already exists", user.Email)
		return errors.New("user with this email already exists")
	}

	log.Printf("User %d registered successfully", user.ID)
	return nil
}

func (r *UserRepoImpl) UpdateUser(user *model.User) error {
	result, err := r.db.Exec(`
		UPDATE users 
		SET username = $1, email = $2, updated_at = NOW()
		WHERE id = $3
	`, user.Username, user.Email, user.ID)

	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check affected rows: %w", err)
	}
	if rowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (r *UserRepoImpl) DeleteUser(id int) error {
	result, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check affected rows: %w", err)
	}
	if rowsAffected == 0 {
		return errors.New("user not found")
	}

	log.Printf("User %d deleted successfully", id)
	return nil
}
