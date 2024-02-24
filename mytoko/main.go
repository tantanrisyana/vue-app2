package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Product struct {
	ID    uint    `json:"id" gorm:"primaryKey"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	// Inisialisasi koneksi ke database MySQL
	dsn := "root:@tcp(127.0.0.1:3306)/shops?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Auto Migrate tabel
	db.AutoMigrate(&User{}, &Product{})

	// Inisialisasi router Gin
	r := gin.Default()
	// Middleware CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"} // Ganti sesuai dengan alamat aplikasi Vue Anda
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	r.Use(cors.New(config))

	// Definisikan endpoint CRUD
	r.GET("/users", GetUsers)
	r.POST("/users", CreateUser)
	r.PUT("/users/:id", UpdateUser)
	r.DELETE("/users/:id", DeleteUser)

	r.GET("/products", GetProducts)
	r.POST("/products", CreateProduct)
	r.PUT("/products/:id", UpdateProduct)
	r.DELETE("/products/:id", DeleteProduct)

	// Jalankan server
	r.Run(":8080")
}

// CRUD handlers for User
func GetUsers(c *gin.Context) {
	users := []User{}
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&newUser)
	c.JSON(http.StatusCreated, newUser)
}

func UpdateUser(c *gin.Context) {
	userID := c.Param("id")
	var updatedUser User

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&User{}).Where("id = ?", userID).Updates(updatedUser)
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUser(c *gin.Context) {
	userID := c.Param("id")
	db.Where("id = ?", userID).Delete(&User{})
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// CRUD handlers for Product
func GetProducts(c *gin.Context) {
	products := []Product{}
	db.Find(&products)
	c.JSON(http.StatusOK, products)
}

func CreateProduct(c *gin.Context) {
	var newProduct Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&newProduct)
	c.JSON(http.StatusCreated, newProduct)
}

func UpdateProduct(c *gin.Context) {
	productID := c.Param("id")
	var updatedProduct Product

	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&Product{}).Where("id = ?", productID).Updates(updatedProduct)
	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

func DeleteProduct(c *gin.Context) {
	productID := c.Param("id")
	db.Where("id = ?", productID).Delete(&Product{})
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
