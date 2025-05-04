package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
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

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server running on port " + port) // untuk menampilkan pesan bahwa server sudah berjalan
	log.Fatal(r.Run(":" + port))                  // untuk menjalankan server pada port yang ditentukan

}
