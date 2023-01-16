package routes

import (
    controller "jwt/controllers"

    "github.com/gin-gonic/gin"
)

//UserRoutes function
func UserRoutes(router *gin.Engine) {
    router.POST("/users/signup", controller.SignUp())
    router.POST("/users/login", controller.Login())
}
