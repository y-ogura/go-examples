package usecase

import (
	"github.com/go-examples/http_sample/domain"
	"github.com/go-examples/http_sample/repository"
)

// userUsecase user usecase
type userUsecase struct {
	userRepo repository.UserRepository
}

// NewUserUsecase mount user usecase
func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: ur,
	}
}

// UserUsecase user usecase interface
type UserUsecase interface {
	List() ([]*domain.User, error)
	Create(user *domain.User) (*domain.User, error)
}

// List get user list usecase
func (u *userUsecase) List() ([]*domain.User, error) {
	res, err := u.userRepo.List()
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Create create new user usecase
func (u *userUsecase) Create(user *domain.User) (*domain.User, error) {
	res, err := u.userRepo.Create(user)
	if err != nil {
		return nil, err
	}
	return res, nil
}
