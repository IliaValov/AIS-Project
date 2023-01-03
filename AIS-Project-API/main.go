package main

import (
	"AIS-Project-API/controllers"
	"AIS-Project-API/database"
	"AIS-Project-API/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {

	println("Starting...")
	database.ConnectDataBase()

	r := gin.Default()

	public := r.Group("/api")

	public.GET("/teacher/grades", controllers.TeacherGrades)
	public.GET("/:subjectId/students", controllers.StudentsPerTeacher)
	public.GET("/student/:studentId/grades", controllers.StudentGrades)
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	r.Run(":8080")

}
