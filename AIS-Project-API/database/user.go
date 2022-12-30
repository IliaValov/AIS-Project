package database

import (
	"AIS-Project-API/utils/token"
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
	Role     string `gorm:"size:255;not null;"`
}

func GetUserByID(uid uint) (User, error) {

	var u User

	if err := DB.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found!")
	}

	u.PrepareGive()

	return u, nil

}

func (u *User) PrepareGive() {
	u.Password = ""
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string) (string, error) {

	var err error

	u := User{}

	err = DB.Model(User{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}

func (u *User) SaveUser() (*User, error) {
	var err error

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return &User{}, err
	}

	u.Password = string(hashedPassword)

	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	// Role field may change when the db relations are finished.
	u.Role = "Default"

	err = DB.Model(&User{}).Create(&u).Error

	if err != nil {
		return &User{}, err
	}

	return u, nil
}

// func (u *User) BeforeSave(tx *gorm.DB) error {

// 	//turn password into hash
// 	println("test")
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return err
// 	}
// 	u.Password = string(hashedPassword)

// 	//remove spaces in username
// 	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

// 	u.Role = "Default"

// 	return nil

// }
