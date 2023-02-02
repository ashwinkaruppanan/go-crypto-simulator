package main

import (
	"log"
	"os"

	"ashwin.com/go-crypto-simulator/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	routes.AuthRoute(router)
	routes.UserRoute(router)

	if err := godotenv.Load(".env"); err != nil {
		log.Panic(err)
	}
	var port string = os.Getenv("PORT")
	router.Run(":" + port)

}
