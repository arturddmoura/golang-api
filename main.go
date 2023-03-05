package main

import (
	"github.com/arturddmoura/golang-api/controllers"
	"github.com/arturddmoura/golang-api/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()

	r.POST("/products", controllers.ProductsCreate)
	r.GET("/products", controllers.GetProducts)
	r.GET("/products/:id", controllers.GetProductById)
	r.PUT("/products/:id", controllers.ProductUpdate)
	r.DELETE("/products/:id", controllers.ProductDelete)

	r.Run()
}
