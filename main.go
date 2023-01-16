package main

import (
	"github.com/gin-gonic/gin"
	routes "jwt/routes"
	config "jwt/config"
	middleware "jwt/middleware"
	controllers "jwt/controllers"
)

func main() {
	config.ConnectDB()
	r := gin.Default()
	routes.UserRoutes(r)

	r.Use(middleware.Authentication())

	// API-2
    r.GET("/users", controllers.GetUsers())

    // API-1
    r.GET("/users/:id", controllers.GetUser())
	
	r.Run(":8080")



}
