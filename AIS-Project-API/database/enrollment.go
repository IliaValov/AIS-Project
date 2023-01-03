package database

import "github.com/jinzhu/gorm"

type Enrollment struct {
	gorm.Model
	StudentId uint64
	Student   Student `gorm:"foreignKey:StudentId;references:UserId;constraint:OnDelete:CASCADE"`
	TeacherId uint64
	Teacher   Teacher `gorm:"foreignKey:TeacherId;references:UserId;constraint:OnDelete:CASCADE"`
	CourseId  uint64
	Course    Course `gorm:"foreignKey:CourseId;references:ID;constraint:OnDelete:CASCADE"`
}
