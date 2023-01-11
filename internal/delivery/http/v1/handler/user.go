package handler

import (
	"go-clean-arch-temp/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	domain.UserUseCase
}

func NewHandler(usecase domain.UserUseCase) Handler {
	return Handler{usecase}
}

func (h *Handler) FetchByUsername() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{})
	}
}

func (h *Handler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusCreated, gin.H{})
	}
}
