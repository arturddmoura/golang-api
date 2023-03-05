package controllers

import (
	"github.com/arturddmoura/golang-api/initializers"
	"github.com/arturddmoura/golang-api/models"
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []models.Product
	initializers.DB.Order("name ASC").Find(&products)

	c.JSON(200, gin.H{"product": products})
}

func GetProductById(c *gin.Context) {
	id := c.Param("id")

	var product models.Product
	result := initializers.DB.First(&product, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"detail": "Product not found"})
		return
	}

	c.JSON(200, gin.H{"product": product})
}

func ProductsCreate(c *gin.Context) {
	var productBody struct {
		Name  string
		Price float32
		Score int
		Image string
	}

	c.Bind(&productBody)

	product := models.Product{Name: productBody.Name, Price: productBody.Price, Score: productBody.Score, Image: productBody.Image}

	result := initializers.DB.Create(&product)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{"product": product})
}

func ProductUpdate(c *gin.Context) {
	id := c.Param("id")

	var productBody struct {
		Name  string
		Price float32
		Score int
		Image string
	}

	c.Bind(&productBody)

	var product models.Product
	result := initializers.DB.First(&product, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"detail": "Product not found"})
		return
	}

	initializers.DB.Model(&product).Updates(models.Product{
		Name: productBody.Name, Price: productBody.Price, Score: productBody.Score, Image: productBody.Image,
	})

	c.JSON(200, gin.H{"product": product})
}

func ProductDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Product{}, id)

	c.JSON(200, gin.H{"product": id})
}
