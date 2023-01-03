package controllers

import (
	"AIS-Project-API/database"
	"AIS-Project-API/utils/token"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GradeInput struct {
	CourseId  uint   `json:"course_id"`
	StudentId uint   `json:"student_id"`
	Grade     uint64 `json:"grade"`
}

func EditGrade(c *gin.Context) {
	var gradeInput GradeInput

	if err := c.ShouldBindJSON(&gradeInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grade := database.Grade{
		StudentId: gradeInput.StudentId,
		CourseId:  gradeInput.CourseId,
		Grade:     gradeInput.Grade,
	}

	grade.Edit()
}

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
