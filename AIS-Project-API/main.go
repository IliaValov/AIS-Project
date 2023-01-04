package main

import (
	"AIS-Project-API/controllers"
	"AIS-Project-API/database"
	"AIS-Project-API/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	println("Starting...")
	database.ConnectDataBase()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	public := r.Group("/api")

	public.GET("/teacher/grades", controllers.TeacherGrades)
	public.GET("/teacher/:teacherId/courses", controllers.CoursesPerTeacher)

	public.GET("/:subjectId/students", controllers.StudentsPerCourse)
	public.GET("/student/:studentId/grades", controllers.StudentGrades)

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/user")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/", controllers.CurrentUser)
	protected.GET("/:id", controllers.GetUserById)

	protectedSubjects := r.Group("/api/subjects")
	protectedSubjects.Use(middlewares.JwtAuthMiddleware())
	protectedSubjects.POST("/editgrade", controllers.EditGrade)
	public.GET("/:subjectId/students", controllers.StudentsPerCourse)
	public.GET("/student/:studentId/grades", controllers.StudentGrades)
	public.GET("/teacher/grades", controllers.TeacherGrades)

	subject := r.Group("/api/subject")
	subject.Use(middlewares.JwtAuthMiddleware())
	subject.POST("joinsubject", controllers.EnrollCourse)

	r.Run(":8080")
}
