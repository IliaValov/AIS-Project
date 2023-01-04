package database

type Grade struct {
	StudentId uint    `gorm:"primaryKey"`
	Student   Student `gorm:"foreignKey:StudentId;references:UserId;constraint:OnDelete:CASCADE"`
	CourseId  uint    `gorm:"primaryKey"`
	Course    Course  `gorm:"foreignKey:CourseId;references:ID;constraint:OnDelete:CASCADE"`
	Grade     uint64  `gorm:"not null;check:grade >= 2 and grade <= 6"`
}

func (g *Grade) Edit() (*Grade, error) {
	var grade Grade
	err := DB.Where("student_id = ? AND course_id = ?", g.StudentId, g.CourseId).Find(&grade).Error
	if err != nil {
		return &Grade{}, err
	}

	if grade == (Grade{}) {
		// has to be created
		err = DB.Create(&g).Error
		if err != nil {
			return &Grade{}, err
		}
	} else {
		// has to be updated
		err = DB.Save(&g).Error
		if err != nil {
			return &Grade{}, err
		}
	}

	return g, nil
}
