package database

import "github.com/jinzhu/gorm"

type Enrollment struct {
	gorm.Model
	StudentId uint
	Student   Student `gorm:"foreignKey:StudentId;references:UserId;constraint:OnDelete:CASCADE"`
	TeacherId uint
	Teacher   Teacher `gorm:"foreignKey:TeacherId;references:UserId;constraint:OnDelete:CASCADE"`
	CourseId  uint
	Course    Course `gorm:"foreignKey:CourseId;references:ID;constraint:OnDelete:CASCADE"`
}

func (e *Enrollment) Enroll() (*Enrollment, error) {
	err := DB.Model(&Enrollment{}).Create(&e).Error
	if err != nil {
		return &Enrollment{}, err
	}

	return e, nil
}
