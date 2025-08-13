package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
	"net/http" 
)

const (
  host     = "localhost"
  port     = 5432
  user     = "admin"
  password = "admin"
  dbname   = "mydb"
)

type Categoria struct {
		Id          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
		Nombre      string `json:"nombre"`
		Descripcion string `json:"descripcion"`
		}


func main() {
	r := gin.Default()

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	fmt.Println("Connected to the database successfully")


	db.AutoMigrate(&Categoria{})

	r.GET("/categorias", func(c *gin.Context) {
		var categorias []Categoria
		if err := db.Find(&categorias).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch categories"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"categorias": categorias})
	})

	r.POST("/categorias", func(c *gin.Context) {
		var newCategory Categoria
		if err := c.ShouldBindJSON(&newCategory); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

        if err := db.Create(&newCategory).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusCreated, gin.H{
            "message":  "Categoría creada",
            "category": newCategory,
        })

	})

	r.DELETE("/categorias/:id", func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&Categoria{}, id).Error; err != nil {
        	c.JSON(http.StatusInternalServerError, gin.H{"error": "Category cannot be deleted"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "Categoría eliminada", "id": id})
	})

	r.Run(":8080")

}
