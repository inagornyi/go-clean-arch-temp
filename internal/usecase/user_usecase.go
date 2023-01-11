package usecase

import "go-clean-arch-temp/internal/domain"

type UserUseCase struct {
	domain.UserRepository
}

func NewUserUseCase(repository domain.UserRepository) UserUseCase {
	return UserUseCase{repository}
}

func (uc UserUseCase) Create(username string) error {
	return nil
}

func (uc UserUseCase) FetchByUsername(username string) (domain.User, error) {
	return domain.User{}, nil
}
