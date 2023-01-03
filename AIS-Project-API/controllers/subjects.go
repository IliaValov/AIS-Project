package controllers

import (
	"AIS-Project-API/database"
	"AIS-Project-API/utils/token"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func TeacherGrades(c *gin.Context) {
	var grades []database.Grade
	database.DB.Find(&grades)

	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := database.GetUserByID(user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var teacherGrades []database.Grade
	for _, currentGrade := range grades {
		if currentGrade.CourseId == user.ID {
			teacherGrades = append(teacherGrades, currentGrade)
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": teacherGrades})
}

func StudentsPerTeacher(c *gin.Context) {
	subjectId, err := strconv.ParseUint(c.Param("subjecId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var enrolments []database.Enrollment
	database.DB.Find(&enrolments)

	var foundStudents []database.Student
	for _, value := range enrolments {
		// TODO change to Course.Teacher.UserId
		if subjectId == value.CourseId {
			foundStudents = append(foundStudents, value.Student)
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": foundStudents})
}

func StudentGrades(c *gin.Context) {
	studentId, err := strconv.ParseUint(c.Param("studentId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var grades []database.Grade
	database.DB.Find(&grades)

	var studentGrades []database.Grade
	for _, currentGrade := range grades {
		if currentGrade.StudentId == uint(studentId) {
			studentGrades = append(studentGrades, currentGrade)
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": studentGrades})
}
