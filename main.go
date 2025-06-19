package main

import (
	"toko_obat/firebase"
	"toko_obat/handlers"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or failed to load .env, relying on system environment variables.")
	}

	// Inisialisasi Firebase App
	firebase.InitFirebase()

	r := gin.Default()

	// Route publik
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to Firebase Auth API with Go!"})
	})

	r.POST("/login", handlers.Login)

	// Route dengan autentikasi Firebase
	api := r.Group("/api")
	api.Use(handlers.AuthMiddleware())
	{
		api.GET("/protected", handlers.ProtectedRoute)
	}

	// Jalankan server
	r.Run(":8080")
}
