package main

import (
	"toko_obat/controller"
	"toko_obat/database"
	"toko_obat/firebase"
	"toko_obat/handlers"
	"toko_obat/repository"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (

	// inits database
	db                 *gorm.DB = database.InitDatabase()
	medicineRepo                = repository.NewMedicineRepository(db)
	medicineController          = controller.NewMedicineController(medicineRepo)

	categoryRepo       = repository.NewCategoryRepository(db)
	categoryController = controller.NewCategoryController(categoryRepo)
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
		api.GET("/medicines", medicineController.GetAllMedicine)
		api.GET("/medicines/:id", medicineController.GetMedicineByID)

		api.GET("/categories", categoryController.GetAllCategory)
		api.GET("/categories/:id", categoryController.GetCategoryByID)
		api.POST("/categories", categoryController.CreateCategory)
	}

	// Jalankan server
	r.Run(":8080")
}
