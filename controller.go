package main

import (
	"Go_server/middleware"
	"Go_server/routes"

	"github.com/gin-gonic/gin"
)

func registerRoutes() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	// r.POST("/registerUser", routes.RegisterUser)
	// r.GET("/getUser", middleware.VerifyAuth, routes.GetUser)
	// r.POST("/update", middleware.VerifyAuth, routes.Update)
	r.POST("/login", routes.Login)
	r.GET("/getbooks", routes.GetBooks)
	r.POST("/addbook", routes.AddBook)
	r.POST("/studentlogin", routes.StudentLogin)
	r.POST("/updatebook", routes.UpdateBook)
	r.POST("./registerstudent", routes.RegisterStudent)
	r.POST("/borrow", routes.Borrow)
	// r.POST("/delete", middleware.VerifyAuth, routes.Delete)
	return r
}
