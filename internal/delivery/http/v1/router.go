package v1

import (
	"go-clean-arch-temp/internal/delivery/http/v1/handler"
	"go-clean-arch-temp/internal/domain"

	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, usecase domain.UserUseCase) {
	v1 := r.Group("/v1")
	{
		user := v1.Group("/user")
		{
			handler := handler.NewHandler(usecase)
			user.GET("/:username", handler.FetchByUsername())
			user.POST("/", handler.Create())
		}
	}
}
