package controllers

import (
	"AIS-Project-API/database"
	"AIS-Project-API/utils/token"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type EnrollInput struct {
	StudentId string `json:"studentId" binding:"required"`
	CourseId  string `json:"courseId" binding:"required"`
}

func EnrollCourse(c *gin.Context) {
	teacherId, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var enrollInput EnrollInput
	if err := c.ShouldBindJSON(&enrollInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	studentId, err := strconv.ParseInt(enrollInput.StudentId, 10, 64)
	courseId, err := strconv.ParseInt(enrollInput.CourseId, 10, 64)

	enrollment := database.Enrollment{
		StudentId: uint(studentId),
		TeacherId: teacherId,
		CourseId:  uint(courseId),
	}

	_, err = enrollment.Enroll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "enrolled successfully",
	})
}
