package handler

import (
	"github.com/gin-gonic/gin"
	mhsDomain "github.com/indraalfauzan/monitoring_skripsi_golang/domain/mahasiswa"
	user "github.com/indraalfauzan/monitoring_skripsi_golang/domain/user"

	"github.com/indraalfauzan/monitoring_skripsi_golang/middleware"
)

func RegisterRoutes(r *gin.Engine, userUC user.UserUseCase, mahasiswaUC mhsDomain.MahasiswaProfileUseCase, profileTAUC mhsDomain.ProfileTAUseCase) {
	authHandler := NewAuthHandler(userUC)
	mahasiswaHandler := NewMahasiswaHandler(mahasiswaUC)
	profileTAHandler := NewProfilTAHandler(profileTAUC)

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
		admin.POST("/add-user", authHandler.RegisterUser)
	}

	mahasiswa := protected.Group("/mahasiswa")
	mahasiswa.Use(middleware.RoleGuard("Mahasiswa"), middleware.AuthMiddleware())
	{
		mahasiswa.POST("/profile", mahasiswaHandler.CreateProfile)
		mahasiswa.GET("/profile", mahasiswaHandler.GetProfile)
		mahasiswa.PUT("/profile", mahasiswaHandler.UpdateProfile)
		mahasiswa.POST("/profile/ta", profileTAHandler.AjukanTA)
		mahasiswa.GET("/profile/ta", profileTAHandler.GetProfileTA)
	}

}
