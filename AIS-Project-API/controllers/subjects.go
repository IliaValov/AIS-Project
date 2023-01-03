package controllers

import (
	"AIS-Project-API/database"
	"github.com/gin-gonic/gin"
	"net/http"
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
