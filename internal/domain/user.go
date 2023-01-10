package domain

import (
	"github.com/google/uuid"
)

type (
	User struct {
		Id   uuid.UUID
		Name string
	}

	UserRepository interface {
		CreateUser(name string) error
		GetUserById(id uuid.UUID) (User, error)
	}

	UserUseCase interface {
		CreateUser(name string) error
		GetUserById(id uuid.UUID) (User, error)
	}
)
