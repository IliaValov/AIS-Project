package main

import (
	"AIS-Project-API/controllers"
	"AIS-Project-API/database"
	"AIS-Project-API/middlewares"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	println("Starting...")
	database.ConnectDataBase()
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/user")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/", controllers.CurrentUser)
	protected.GET("/:id", controllers.GetUserById)

	protectedSubjects := r.Group("/api/subjects")
	protectedSubjects.Use(middlewares.JwtAuthMiddleware())
	protectedSubjects.POST("/editgrade", controllers.EditGrade)
	protectedSubjects.GET("/:subjectId/students", controllers.StudentsPerCourse)
	protectedSubjects.GET("/student/:studentId/grades", controllers.StudentGrades)
	protectedSubjects.GET("/teacher/:teacherId/courses", controllers.CoursesPerTeacher)
	protectedSubjects.GET("/teacher/course/:courseId/students/grades", controllers.StudentsAndGradesPerCourses)
	protectedSubjects.POST("/joinsubject", controllers.EnrollCourse)
	protectedSubjects.GET("/:subjectId/students/notenrolled", controllers.NotEnrolled)

	r.RunTLS(":8080", os.Getenv("CERT_PATH"), os.Getenv("CERT_PRIVATE_KEY_PAHT"))
}
