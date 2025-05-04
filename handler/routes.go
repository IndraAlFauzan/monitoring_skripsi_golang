package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/indraalfauzan/monitoring_skripsi_golang/domain"
)

func RegisterRoutes(r *gin.Engine, userUC domain.UserUseCase) {
	authHandler := NewAuthHandler(userUC)

	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}
}
