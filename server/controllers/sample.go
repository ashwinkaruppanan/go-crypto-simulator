package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"ashwin.com/go-crypto-simulator/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

var collection = database.OpenCollection(database.Client, "sample")

func Sample() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

		inserted, err := collection.InsertOne(ctx, bson.M{"name": "ranjith"})
		if err != nil {
			log.Panic(err)
		}
		defer cancel()
		c.JSON(http.StatusOK, inserted)
	}
}
