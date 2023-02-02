package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"ashwin.com/go-crypto-simulator/database"
	"ashwin.com/go-crypto-simulator/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginDetails *models.Users
		var dbDetails any

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		if bindErr := c.ShouldBindJSON(&loginDetails); bindErr != nil {
			log.Panic(bindErr)
		}
		opts := options.FindOne().SetProjection(bson.M{
			"_id":      1,
			"name":     1,
			"password": 1,
		})
		dbErr := userCollection.FindOne(ctx, bson.M{"email": loginDetails.Email}, opts).Decode(&dbDetails)
		defer cancel()
		if dbErr != nil {
			log.Panic(dbErr)
		}

	}
}

func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("token", "", -1, "/", "", false, true)
		c.JSON(http.StatusOK, "Logged out successfully")
	}
}
