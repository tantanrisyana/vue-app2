// main.go
package main

import (
	"log"
	"net/http"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var db *gorm.DB

type Task struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Task string `json:"task"`
	Time string `json:"time"`
}

func initDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true"

	var errDB error
	db, errDB = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errDB != nil {
		log.Fatal(errDB)
	}

	db.AutoMigrate(&Task{})
}

func getTasks(c *gin.Context) {
	var tasks []Task
	db.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

func addTask(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&task)
	c.JSON(http.StatusOK, task)
}

func editTask(c *gin.Context) {
	id := c.Params.ByName("id")
	var task Task

	if err := db.Where("id = ?", id).First(&task).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.BindJSON(&task)
	db.Save(&task)
	c.JSON(http.StatusOK, task)
}

func deleteTask(c *gin.Context) {
	id := c.Params.ByName("id")
	var task Task

	if err := db.Where("id = ?", id).First(&task).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	db.Delete(&task)
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func main() {
	initDB()

	r := gin.Default()

	// CORS middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	r.GET("/tasks", getTasks)
	r.POST("/tasks", addTask)
	r.PUT("/tasks/:id", editTask)
	r.DELETE("/tasks/:id", deleteTask)

	r.Run(":8080")
}
