package main

import (
	"github.com/arturddmoura/golang-api/initializers"
	"github.com/arturddmoura/golang-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Product{})
}
