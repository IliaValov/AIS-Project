package main

import (
	"AIS-Project-API/controllers"
	"AIS-Project-API/database"
	"AIS-Project-API/middlewares"

	"github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func main() {

	println("Starting...")
	database.ConnectDataBase()

	r := gin.Default()

    r.Use(cors.New(cors.Config{
        AllowOrigins: []string{"http://localhost:4200"},
        AllowMethods: []string{"GET", "POST"},
        AllowHeaders: []string{"*"},
        ExposeHeaders: []string{"*"},
        AllowCredentials: true,
    })) 

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	protectedSubjects := r.Group("/api/subjects")
	protectedSubjects.Use(middlewares.JwtAuthMiddleware())
	protectedSubjects.POST("/editgrade", controllers.EditGrade)

	subject := r.Group("/api/subject")
	subject.Use(middlewares.JwtAuthMiddleware())
	subject.POST("joinsubject", controllers.EnrollCourse)

	r.Run(":8080")
}
