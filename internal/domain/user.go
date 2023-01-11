package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	User struct {
		Id       uuid.UUID `json:"id"`
		Username string    `json:"username"`
	}

	UserEntity struct {
		ID       string
		Username string
	}

	UserRepository interface {
		CreateUser(username string) error
		FetchByUsername(username string) (UserEntity, error)
	}

	UserUseCase interface {
		Create(username string) error
		FetchByUsername(username string) (User, error)
	}

	UserHandler interface {
		Create() gin.HandlerFunc
		FetchByUsername() gin.HandlerFunc
	}
)
