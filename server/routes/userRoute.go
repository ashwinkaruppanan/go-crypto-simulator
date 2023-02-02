package routes

import (
	"ashwin.com/go-crypto-simulator/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/api/v1/balance", controllers.GetBalance())
}
