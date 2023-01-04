package database

type Enrollment struct {
	StudentId uint    `gorm:"primaryKey"`
	Student   Student `gorm:"foreignKey:StudentId;references:UserId;constraint:OnDelete:CASCADE"`
	TeacherId uint    `gorm:"primaryKey"`
	Teacher   Teacher `gorm:"foreignKey:TeacherId;references:UserId;constraint:OnDelete:CASCADE"`
	CourseId  uint    `gorm:"primaryKey"`
	Course    Course  `gorm:"foreignKey:CourseId;references:ID;constraint:OnDelete:CASCADE"`
}

func (e *Enrollment) Enroll() (*Enrollment, error) {
	err := DB.Model(&Enrollment{}).Create(&e).Error
	if err != nil {
		return &Enrollment{}, err
	}

	return e, nil
}
