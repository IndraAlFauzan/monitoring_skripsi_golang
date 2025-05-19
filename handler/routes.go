package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/indraalfauzan/monitoring_skripsi_golang/domain"

	"github.com/indraalfauzan/monitoring_skripsi_golang/middleware"
)

func RegisterRoutes(r *gin.Engine, userUC domain.UserUseCase, mahasiswaUC domain.MahasiswaProfileUseCase) {
	authHandler := NewAuthHandler(userUC)
	mahasiswaHandler := NewMahasiswaHandler(mahasiswaUC)

	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.RegisterMhs)
		auth.POST("/login", authHandler.Login)
	}

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/me", GetProfile)
	}

	admin := protected.Group("/admin")
	admin.Use(middleware.RoleGuard("Admin"), middleware.AuthMiddleware())
	{
		admin.POST("/register", authHandler.RegisterUser)
	}

	mahasiswa := protected.Group("/mahasiswa")
	mahasiswa.Use(middleware.RoleGuard("Mahasiswa"), middleware.AuthMiddleware())
	{
		mahasiswa.POST("/profile", mahasiswaHandler.CreateProfile)
		mahasiswa.GET("/profile", mahasiswaHandler.GetProfile)
		mahasiswa.PUT("/profile", mahasiswaHandler.UpdateProfile)
	}

}
