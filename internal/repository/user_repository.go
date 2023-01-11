package repository

import "go-clean-arch-temp/internal/domain"

type UserRepository struct{}

func NewUserRepository() UserRepository {
	return UserRepository{}
}

func (r UserRepository) CreateUser(username string) error {
	return nil
}

func (r UserRepository) FetchByUsername(username string) (domain.UserEntity, error) {
	return domain.UserEntity{}, nil
}
