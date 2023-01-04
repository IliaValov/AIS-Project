package controllers

import (
	"AIS-Project-API/database"
	"AIS-Project-API/services"
	"AIS-Project-API/utils/token"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserById(c *gin.Context) {
	uid, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	adminRights, err := token.ExtractAdminRights(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !adminRights {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	userId := uint(uid)

	u, err := database.GetUserByID(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

func CurrentUser(c *gin.Context) {
	userId, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": userId})
}

type LoginInput struct {
	Username string `json:"username" binding:"required,lte=20"`
	Password string `json:"password" binding:"required,lte=20"`
}

func Login(c *gin.Context) {

	var input LoginInput

	err := services.ValidateInput(c, &input)
	if err != nil {
		return
	}

	u := database.User{}

	u.Username = input.Username
	u.Password = input.Password

	token, err := database.LoginCheck(u.Username, u.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

type RegisterInput struct {
	FirstName string `json:"firstName" binding:"required,lte=20"`
	LastName  string `json:"lastName" binding:"required,lte=20"`
	Password  string `json:"password" binding:"required,gte=6,lte=20"`
}

func Register(c *gin.Context) {
	var input RegisterInput

	err := services.ValidateInput(c, &input)
	if err != nil {
		return
	}

	u := database.User{
		Password:    input.Password,
		AdminRights: false,
	}
	user, err := u.SaveUser(input.FirstName, input.LastName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "registration success",
		"username": user.Username,
	})
}
