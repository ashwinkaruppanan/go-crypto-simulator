package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"ashwin.com/go-crypto-simulator/database"
	"ashwin.com/go-crypto-simulator/helper"
	"ashwin.com/go-crypto-simulator/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser *models.Users

		if bindErr := c.ShouldBindJSON(&newUser); bindErr != nil {
			log.Panic(bindErr)
		}

		if validationErr := validator.New().Struct(newUser); validationErr != nil {
			c.JSON(http.StatusBadRequest, validationErr)
			return
		}

		fmt.Println(newUser)
		newUser.ID = primitive.NewObjectID()
		newUser.Password = helper.HashPassword(newUser.Password)
		newUser.Fiat = 10000.0
		newUser.Bitcoin = 0.0009
		newUser.JoinedAt = time.Now().Unix()
		newUser.UpdatedAt = time.Now().Unix()

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		count, dbErr := userCollection.CountDocuments(ctx, bson.M{"email": newUser.Email})
		defer cancel()
		if dbErr != nil {
			log.Panic(dbErr)
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered with some other account"})
			return
		}

		_, insertedErr := userCollection.InsertOne(ctx, newUser)
		defer cancel()

		if insertedErr != nil {
			log.Panic(insertedErr)
		}

		c.JSON(http.StatusOK, gin.H{"success": "registration successful"})

	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {

		var loginDetails *models.Users
		var dbDetails *models.Users

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		if bindErr := c.ShouldBindJSON(&loginDetails); bindErr != nil {
			log.Panic(bindErr)
		}
		opts := options.FindOne().SetProjection(bson.M{
			"_id":       1,
			"full_name": 1,
			"password":  1,
		})
		dbErr := userCollection.FindOne(ctx, bson.M{"email": loginDetails.Email}, opts).Decode(&dbDetails)
		defer cancel()
		if dbErr != nil {
			// log.Panic(dbErr)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Email"})
			return
		}

		checkPassword := helper.VerifyPassword(loginDetails.Password, dbDetails.Password)

		if !checkPassword {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Password"})
			return
		}

		token := helper.GenerateToken(dbDetails.Name, dbDetails.ID.Hex())

		c.SetCookie("token", token, 60*60*3, "/", "localhost", false, true)
		c.JSON(http.StatusOK, "Login Successfull")
	}
}

func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("token", "", -1, "/", "", false, true)
		c.JSON(http.StatusOK, "Logged out successfully")
	}
}
