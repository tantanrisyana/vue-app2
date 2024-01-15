package main

import (
	"fmt"
	"tugas/controllers"
	"tugas/models"
	"tugas/utils"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	configs := utils.LoadConfigs()

	db, err := utils.InitDB(configs.DBUsername, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName)
	if err != nil {
		fmt.Printf("Failed to connect to the database: %v\n", err)
		return
	}

	db.AutoMigrate(&models.Siswa{})

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())

	// Middleware CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"}, // Sesuaikan dengan alamat aplikasi Vue Anda
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))

	siswaController := controllers.NewSiswaController(db)

	// Routes
	e.GET("/siswa", siswaController.GetSiswaList)
	e.GET("/siswa/:id", siswaController.GetSiswaByID)
	e.POST("/siswa", siswaController.CreateSiswa)
	e.PUT("/siswa/:id", siswaController.UpdateSiswa)
	e.DELETE("/siswa/:id", siswaController.DeleteSiswa)

	// Start server
	e.Start(":8080")
}
