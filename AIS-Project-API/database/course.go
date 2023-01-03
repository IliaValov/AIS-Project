package database

import "github.com/jinzhu/gorm"

type Course struct {
	gorm.Model
	Name string `gorm:"size:255;not null"`
}
