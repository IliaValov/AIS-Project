package database

type Teacher struct {
	UserId uint   `gorm:"unique"`
	User   User   `gorm:"foreignKey:UserId;references:ID;constraint:OnDelete:CASCADE"`
	Name   string `gorm:"size:255;not null"`
}
