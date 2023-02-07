package routes

import (
	"net/http"

	"ashwin.com/go-crypto-simulator/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoute(router *gin.Engine) {
	router.GET("/", root())
	router.POST("/api/v1/login/", controllers.Login())
	router.POST("/api/v1/signup/", controllers.Signup())
	router.DELETE("/api/v1/logout/", controllers.Logout())
}

func root() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.SetCookie("token", "mystring", 3600, "/", "localhost", false, true)
		ctx.JSON(http.StatusOK, "Home")
	}
}
