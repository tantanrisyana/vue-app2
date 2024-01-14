package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type DataModel struct {
	gorm.Model
	Name     string
	ParentID *uint
}

func main() {
	r := gin.Default()

	initDB()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"} // Ganti dengan alamat frontend Vue Anda
	r.Use(cors.New(config))

	r.GET("/api/data", getData)

	r.Run(":3000")
}

func initDB() {
	var err error
	db, err = gorm.Open("mysql", "root:@tcp(localhost:3306)/userdb?parseTime=true")
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&DataModel{})
}

func getData(c *gin.Context) {
	var data []DataModel
	if err := db.Find(&data).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch data"})
		return
	}

	c.JSON(200, data)
}
