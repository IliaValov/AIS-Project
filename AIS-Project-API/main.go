package main

import (
	"AIS-Project-API/controllers"
	"AIS-Project-API/database"
	"AIS-Project-API/middlewares"

	"github.com/gin-gonic/gin"
)

//[POST] api/subjects/editgrade
//{
//"studentId": 1,
//"subjectId": 1
//“grade” : 2
//}
func main() {

	println("Starting...")
	database.ConnectDataBase()

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	protectedSubjects := r.Group("/api/subjects")
	protectedSubjects.Use(middlewares.JwtAuthMiddleware())
	protectedSubjects.POST("/editgrade", controllers.EditGrade)

	r.Run(":8080")

}
