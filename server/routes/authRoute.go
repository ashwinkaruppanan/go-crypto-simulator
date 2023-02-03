package routes

import (
	"ashwin.com/go-crypto-simulator/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoute(router *gin.Engine) {
	router.POST("/api/v1/login/", controllers.Login())
	router.POST("/api/v1/signup/", controllers.Signup())
	router.DELETE("/api/v1/logout/", controllers.Logout())
}
