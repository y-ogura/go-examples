package repository

import (
	"database/sql"

	"github.com/go-examples/http_sample/domain"
	uuid "github.com/satori/go.uuid"
)

// userRepository user repository
type userRepository struct {
	Conn *sql.DB
}

// NewUserRepository mount user repository
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		Conn: db,
	}
}

// UserRepository user repository interface
type UserRepository interface {
	List() ([]*domain.User, error)
	Create(user *domain.User) (*domain.User, error)
}

// List get user list
func (u *userRepository) List() ([]*domain.User, error) {
	var res []*domain.User
	user := &domain.User{
		ID:       uuid.NewV4(),
		Email:    "example@example.com",
		Password: "abc12345",
	}
	res = append(res, user)
	return res, nil
}

// Create create new user
func (u *userRepository) Create(user *domain.User) (*domain.User, error) {
	return nil, nil
}
