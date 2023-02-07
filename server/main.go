package main

import (
	"log"
	"net/http"
	"os"

	"ashwin.com/go-crypto-simulator/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Next()
	})

	router.OPTIONS("/api/v1/login/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	routes.AuthRoute(router)
	routes.UserRoute(router)
	//go helper.Timer()

	if err := godotenv.Load(".env"); err != nil {
		log.Panic(err)
	}
	var port string = os.Getenv("PORT")
	router.Run(":" + port)

}
