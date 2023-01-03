package database

import "fmt"

type Grade struct {
	StudentId uint    `gorm:"primaryKey"`
	Student   Student `gorm:"foreignKey:StudentId;references:UserId;constraint:OnDelete:CASCADE"`
	CourseId  uint    `gorm:"primaryKey"`
	Course    Course  `gorm:"foreignKey:CourseId;references:ID;constraint:OnDelete:CASCADE"`
	Grade     uint64  `gorm:"not null;check:grade >= 2 and grade <= 6"`
}

func (g *Grade) Edit() {
	var grade Grade
	DB.Where("student_id = ? AND course_id = ?", g.StudentId, g.CourseId).Find(&grade)

	if grade == (Grade{}) {
		// has to be created
		fmt.Println("CRETE GRADE", g)
		DB.Create(&g)
	} else {
		// has to be updated
		DB.Update("grade", g.Grade)
	}
}
