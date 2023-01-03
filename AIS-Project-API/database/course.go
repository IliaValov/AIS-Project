package database

import "github.com/jinzhu/gorm"

type Course struct {
	gorm.Model
	Name      string  `gorm:"size:255;not null"`
	TeacherId uint    `gorm:"unique"`
	Teacher   Teacher `gorm:"foreignKey:TeacherId;references:UserId;constraint:OnDelete:CASCADE"`
}
