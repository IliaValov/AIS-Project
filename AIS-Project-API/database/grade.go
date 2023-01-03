package database

import "github.com/jinzhu/gorm"

type Grade struct {
	gorm.Model
	StudentId uint
	Student   Student `gorm:"foreignKey:StudentId;references:UserId;constraint:OnDelete:CASCADE"`
	CourseId  uint
	Course    Course `gorm:"foreignKey:CourseId;references:ID;constraint:OnDelete:CASCADE"`
	Grade     uint64 `gorm:"not null;check:grade >= 2 and grade <= 6"`
}
