package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID    int
	Name  string
	Price float32
	Score int
	Image string
}
