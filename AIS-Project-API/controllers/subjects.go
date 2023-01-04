package controllers

import (
	"AIS-Project-API/database"
	"AIS-Project-API/services"
	"AIS-Project-API/utils/token"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GradeInput struct {
	CourseId  uint   `json:"courseId" binding:"required,gte=1"`
	StudentId uint   `json:"studentId" binding:"required,gte=1"`
	Grade     uint64 `json:"grade" binding:"required,gte=2,lte=6"`
}

func EditGrade(c *gin.Context) {
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

	var gradeInput GradeInput

	if services.ValidateInput(c, &gradeInput) != nil {
		return
	}

	grade := database.Grade{
		StudentId: gradeInput.StudentId,
		CourseId:  gradeInput.CourseId,
		Grade:     gradeInput.Grade,
	}

	teacherId, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var enrollment database.Enrollment
	if err := database.DB.Where("teacherId = ? AND courseId = ? AND studentId = ?",
		teacherId, grade.CourseId, grade.StudentId).Find(&enrollment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if enrollment == (database.Enrollment{}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "teacher is unauthorized to edit this grade",
		})
		return
	}

	_, err = grade.Edit()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "grade edited successfully",
	})
}

type EnrollInput struct {
	StudentId string `json:"studentId" binding:"required,gte=1"`
	CourseId  string `json:"courseId" binding:"required,gte=1"`
}

func EnrollCourse(c *gin.Context) {
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

	teacherId, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var enrollInput EnrollInput
	if services.ValidateInput(c, &enrollInput) != nil {
		return
	}

	studentId, err := strconv.ParseInt(enrollInput.StudentId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	courseId, err := strconv.ParseInt(enrollInput.CourseId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

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

// func TeacherGrades(c *gin.Context) {
// 	adminRights, err := token.ExtractAdminRights(c)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if !adminRights {
// 		c.JSON(http.StatusForbidden, gin.H{
// 			"error": "Unauthorized",
// 		})
// 		return
// 	}

// 	user_id, err := token.ExtractTokenID(c)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	var enrollment []database.Enrollment
// 	database.DB.Where("teacher_id = ?", user_id).Find(&enrollment)

// 	type ResultType struct {
// 		Student database.Student
// 		Grade   database.Grade
// 	}

// 	var grades []database.Grade
// 	for _, value := range enrollment {
// 		var tempGrades []database.Grade
// 		database.DB.Where("student_id = ?", value.StudentId).Find(&tempGrades)
// 		grades = append(grades, tempGrades...)
// 	}

// 	fmt.Println(grades)
// 	c.JSON(http.StatusOK, gin.H{"data": grades})

// 	// var grades []database.Grade
// 	// database.DB.Find(&grades)

// 	// var teacherGrades []database.Grade
// 	// for _, currentGrade := range grades {
// 	// 	if currentGrade.CourseId == user.ID {
// 	// 		teacherGrades = append(teacherGrades, currentGrade)
// 	// 	}
// 	// }

// 	// c.JSON(http.StatusOK, gin.H{"data": teacherGrades})
// }

// за учител връща неговите курсове?
func CoursesPerTeacher(c *gin.Context) {
	teacherId, err := strconv.ParseUint(c.Param("teacherId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error while parsing parameter: ": err.Error()})
		return
	}

	var courses []database.Course
	if err := database.DB.Preload("Teacher").Where("teacher_id = ?", teacherId).Find(&courses).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error while getting courses from database: ": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": courses})
}

// for teacher's course return students and grades
func StudentsAndGradesPerCourses(c *gin.Context) {
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

	courseId, err := strconv.ParseUint(c.Param("courseId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error while parsing parameter: ": err.Error()})
		return
	}

	teacherId, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	type StudentGrade struct {
		Student database.Student
		Grade   database.Grade
	}

	var enrollments []database.Enrollment
	if err := database.DB.Where("teacher_id = ? AND course_id = ?", teacherId, courseId).Find(&enrollments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error while getting enrollments: ": err})
		return
	}

	var grades []database.Grade
	if err := database.DB.Where("course_id = ?", courseId).Find(&grades).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error while getting grades: ": err})
		return
	}

	var result []StudentGrade
	for _, currentGrade := range grades {
		for _, currentEnrollment := range enrollments {
			if currentEnrollment.StudentId == currentGrade.StudentId {
				result = append(result, StudentGrade{
					currentGrade.Student, currentGrade,
				})
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

// Returns array of Students that are enrolled in the subject passed in the url
func StudentsPerCourse(c *gin.Context) {
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

	subjectId, err := strconv.ParseUint(c.Param("subjectId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error while parsing parameter: ": err})
		return
	}

	var enrollments []database.Enrollment
	if err := database.DB.Where("course_id = ?", subjectId).Find(&enrollments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error while getting enrollments: ": err})
		return
	}

	var students []database.Student
	if err := database.DB.Find(&students).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error while getting students: ": err})
		return
	}

	var foundStudents []database.Student
	for _, currentEnrollments := range enrollments {
		for _, currentStudent := range students {
			if currentStudent.UserId == uint(currentEnrollments.StudentId) {
				foundStudents = append(foundStudents, currentStudent)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": foundStudents})
}

// Returns array of ResultGrade - tuple of CourseName and GradeNumber
func StudentGrades(c *gin.Context) {
	studentId, err := strconv.ParseUint(c.Param("studentId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error while parsing parameter: ": err.Error()})
		return
	}

	var grades []database.Grade
	if err := database.DB.Where("student_id = ?", studentId).Find(&grades).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error while getting grades: ": err})
		return
	}

	var courses []database.Course
	if err := database.DB.Find(&courses).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error while getting courses: ": err})
		return
	}

	type ResultGrade struct {
		CourseName  string
		GradeNumber uint
	}

	var result []ResultGrade
	for _, currentGrade := range grades {
		for _, currentCourse := range courses {
			if currentCourse.ID == currentGrade.CourseId {
				result = append(result, ResultGrade{
					currentCourse.Name, uint(currentGrade.Grade),
				})
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}
