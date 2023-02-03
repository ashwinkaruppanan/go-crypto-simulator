package routes

import (
	"ashwin.com/go-crypto-simulator/controllers"
	"ashwin.com/go-crypto-simulator/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.Use(middleware.AuthenticateMiddleware())
	router.GET("/api/v1/balance/", controllers.GetBalance())
	router.GET("/api/v1/open/", controllers.OpenOrders())
	router.GET("/api/v1/history/", controllers.TradeHistory())
	router.POST("/api/v1/limit/buy/", controllers.LmitBuy())
	router.POST("/api/v1/limit/sell/", controllers.LimitSell())
	router.POST("/api/v1/market/buy/", controllers.MarketBuy())
	router.POST("/api/v1/market/sell/", controllers.MarketSell())
	//router.DELETE("/api/v1/cancel/all/", controllers.CancelAllOpenOrders())
	router.DELETE("/api/v1/cancel/", controllers.CancelOrderById())
}
