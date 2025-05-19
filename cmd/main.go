package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/indraalfauzan/monitoring_skripsi_golang/config"

	"github.com/indraalfauzan/monitoring_skripsi_golang/handler"
	"github.com/indraalfauzan/monitoring_skripsi_golang/repository"
	"github.com/indraalfauzan/monitoring_skripsi_golang/usecase"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set Gin mode based on ENV
	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = gin.ReleaseMode // default to release if not set
	}
	gin.SetMode(mode)

	r := gin.Default() // untuk membuat router baru

	// Set trusted proxies explicitly (safer)
	if err := r.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatal("Failed to set trusted proxies:", err)
	}
	// Connect DB
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	// // Auto migrate table
	// _ = db.AutoMigrate(&entity.User{}, &entity.Role{})

	// DI wiring
	userRepo := repository.NewUserRepository(db)
	userUC := usecase.NewUserUseCase(userRepo)

	mhsRepo := repository.NewMahasiswaProfileRepository(db)
	mhsUC := usecase.NewMahasiswaProfileUsecase(mhsRepo)

	// Register all routes
	handler.RegisterRoutes(r, userUC, mhsUC)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server running on port " + port) // untuk menampilkan pesan bahwa server sudah berjalan
	log.Fatal(r.Run(":" + port))                  // untuk menjalankan server pada port yang ditentukan

}
