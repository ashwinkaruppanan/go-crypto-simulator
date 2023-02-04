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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ordersCollection *mongo.Collection = database.OpenCollection(database.Client, "orders")

func GetBalance() gin.HandlerFunc {
	return func(c *gin.Context) {
		//getting user id from request context
		ID := c.GetString("id")
		userID, idErr := primitive.ObjectIDFromHex(ID)
		if idErr != nil {
			log.Panic(idErr)
		}

		//getting balance of user in db
		var dbDetails *models.Users
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		opts := options.FindOne().SetProjection(bson.M{
			"fiat":    1,
			"bitcoin": 1,
		})
		dbErr := userCollection.FindOne(ctx, bson.M{"_id": userID}, opts).Decode(&dbDetails)
		defer cancel()
		if dbErr != nil {
			log.Panic(dbErr)
		}

		//sending response
		c.JSON(http.StatusOK, gin.H{
			"fiat":    dbDetails.Fiat,
			"bitcoin": dbDetails.Bitcoin,
		})
	}
}

func OpenOrders() gin.HandlerFunc {
	return func(c *gin.Context) {

		//getting user id from request context
		id := c.GetString("id")
		userID, idErr := primitive.ObjectIDFromHex(id)
		if idErr != nil {
			log.Panic(idErr)
		}

		//getting user's open orders from db collection
		var dbDetails []*models.Orders
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		cursor, dbErr := ordersCollection.Find(ctx, bson.M{"user_id": userID, "status": "OPEN"})
		defer cancel()
		if dbErr != nil {
			log.Panic(dbErr)
			c.JSON(http.StatusBadRequest, gin.H{"error": dbErr})
		}

		if err := cursor.All(ctx, &dbDetails); err != nil {
			log.Panic(err)
		}

		//sendin response
		c.JSON(http.StatusOK, dbDetails)
	}
}

func TradeHistory() gin.HandlerFunc {
	return func(c *gin.Context) {
		//getting user id from request context
		id := c.GetString("id")
		userID, idErr := primitive.ObjectIDFromHex(id)
		if idErr != nil {
			log.Panic(idErr)
		}

		//getting user's Trade History from db collection
		var dbDetails []*models.Orders
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		cursor, dbErr := ordersCollection.Find(ctx, bson.M{"user_id": userID, "status": "EXECUTED"})
		defer cancel()
		if dbErr != nil {
			log.Panic(dbErr)
			c.JSON(http.StatusBadRequest, gin.H{"error": dbErr})
		}
		if err := cursor.All(ctx, &dbDetails); err != nil {
			log.Panic(err)
		}

		//sending response
		c.JSON(http.StatusOK, dbDetails)
	}
}

// strct for getting btc price and total trade amount from the client

func LmitBuy() gin.HandlerFunc {
	return func(c *gin.Context) {
		type Limit struct {
			BTCprice float32 `json:"price"`
			TotalUSD float32 `json:"total"`
		}

		var newBuyOrder *Limit = new(Limit)
		var newOpenOrder *models.Orders = new(models.Orders)
		var dbBalance *models.Users = new(models.Users)

		//bind json from the request
		if bindErr := c.ShouldBindJSON(&newBuyOrder); bindErr != nil {
			log.Panic(bindErr)
		}

		//getting user id form the request context
		id := c.GetString("id")
		userID, idErr := primitive.ObjectIDFromHex(id)
		if idErr != nil {
			log.Panic(idErr)
		}

		//checking fiat balance before taking buy order
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		opts := options.FindOne().SetProjection(bson.M{
			"fiat": 1,
		})
		dbErr := userCollection.FindOne(ctx, bson.M{"_id": userID}, opts).Decode(&dbBalance)
		defer cancel()
		if dbErr != nil {
			log.Panic(dbErr)
		}

		if dbBalance.Fiat < 10.01 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient USD"})
			return
		}
		if newBuyOrder.TotalUSD > dbBalance.Fiat {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient USD"})
			return
		}

		//assigning new values for buy order
		newOpenOrder.OrderID = primitive.NewObjectID()
		newOpenOrder.UserID = userID
		newOpenOrder.Status = "OPEN"
		newOpenOrder.Pair = "BTC/USD"
		newOpenOrder.Type = "LIMIT"
		newOpenOrder.Side = "BUY"
		newOpenOrder.Price = newBuyOrder.BTCprice
		newOpenOrder.Amount = newBuyOrder.TotalUSD / newBuyOrder.BTCprice
		newOpenOrder.Total = newBuyOrder.TotalUSD
		newOpenOrder.OpenedAt = time.Now().Unix()

		//updating open order collection
		_, insertErr := ordersCollection.InsertOne(ctx, newOpenOrder)
		defer cancel()
		if insertErr != nil {
			log.Panic(insertErr)
		}

		//updating user's fiat balance
		filter := bson.M{"_id": userID}
		update := bson.M{"$set": bson.M{
			"fiat": dbBalance.Fiat - newBuyOrder.TotalUSD,
		}}
		_, updateErr := userCollection.UpdateOne(ctx, filter, update)
		defer cancel()

		if updateErr != nil {
			log.Panic(updateErr)
		}

		//seding response
		c.JSON(http.StatusOK, gin.H{"success": fmt.Sprint("Buy order at ", newBuyOrder.BTCprice)})
	}
}

func LimitSell() gin.HandlerFunc {
	return func(c *gin.Context) {
		type Limit struct {
			BTCprice  float32 `json:"price"`
			BTCamount float32 `json:"amount"`
		}

		var newSellOrder *Limit = new(Limit)
		var newOpenOrder *models.Orders = new(models.Orders)
		var dbBalance *models.Users = new(models.Users)

		//bind json from request
		if bindErr := c.ShouldBindJSON(&newSellOrder); bindErr != nil {
			log.Panic(bindErr)
		}

		//getting user id from request context
		id := c.GetString("id")
		userID, idErr := primitive.ObjectIDFromHex(id)
		if idErr != nil {
			log.Panic(idErr)
		}

		//checking btc quantity before taking sell order
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		opts := options.FindOne().SetProjection(bson.M{
			"bitcoin": 1,
		})
		dbErr := userCollection.FindOne(ctx, bson.M{"_id": userID}, opts).Decode(&dbBalance)
		defer cancel()
		if dbErr != nil {
			log.Panic(dbErr)
		}

		if dbBalance.Bitcoin < 0.0004 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient BTC"})
			return
		}

		if newSellOrder.BTCamount > dbBalance.Bitcoin {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient BTC"})
			return
		}

		//assigning new values for sell order
		newOpenOrder.OrderID = primitive.NewObjectID()
		newOpenOrder.UserID = userID
		newOpenOrder.Status = "OPEN"
		newOpenOrder.Pair = "BTC/USD"
		newOpenOrder.Type = "LIMIT"
		newOpenOrder.Side = "SELL"
		newOpenOrder.Price = newSellOrder.BTCprice
		newOpenOrder.Amount = newSellOrder.BTCamount
		newOpenOrder.Total = newSellOrder.BTCamount * newSellOrder.BTCprice
		newOpenOrder.OpenedAt = time.Now().Unix()

		//updating open order collection
		_, insertErr := ordersCollection.InsertOne(ctx, newOpenOrder)
		defer cancel()
		if insertErr != nil {
			log.Panic(insertErr)
		}
		filter := bson.M{"_id": userID}
		update := bson.M{"$set": bson.M{
			"bitcoin": dbBalance.Bitcoin - newOpenOrder.Amount,
		}}

		//updating btc quantity in user collection - blocking open order btc in user collection
		_, updateErr := userCollection.UpdateOne(ctx, filter, update)
		defer cancel()

		if updateErr != nil {
			log.Panic(updateErr)
		}

		//sending response
		c.JSON(http.StatusOK, gin.H{"success": fmt.Sprint("Sell order at ", newSellOrder.BTCprice)})
	}
}

func MarketBuy() gin.HandlerFunc {
	return func(c *gin.Context) {
		type Market struct {
			TotalUSD float32 `json:"total"`
		}

		var totalUSD *Market = new(Market)

		// bind json from request
		if bindErr := c.ShouldBindJSON(&totalUSD); bindErr != nil {
			log.Panic(bindErr)
		}

		//getting user id from request context
		id := c.GetString("id")
		userID, idErr := primitive.ObjectIDFromHex(id)
		if idErr != nil {
			log.Panic(idErr)
		}

		// checking user's balance from db collection
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		var dbBalance *models.Users = new(models.Users)
		opts := options.FindOne().SetProjection(bson.M{"fiat": 1, "bitcoin": 1})
		dbErr := userCollection.FindOne(ctx, bson.M{"_id": userID}, opts).Decode(&dbBalance)
		defer cancel()
		if dbErr != nil {
			log.Panic(dbErr)
		}

		if dbBalance.Fiat < 10.01 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient USD"})
			return
		}
		if dbBalance.Fiat < totalUSD.TotalUSD {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient USD"})
			return
		}

		// getting current market price with binance api
		marketPrice, priceErr := helper.CurrentBTCPrice()
		if priceErr != nil {
			log.Panic(priceErr)
		}

		//updating balance in user collection
		_, updateErr := userCollection.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$set": bson.M{
			"fiat":    dbBalance.Fiat - totalUSD.TotalUSD,
			"bitcoin": dbBalance.Bitcoin + totalUSD.TotalUSD/marketPrice,
		}})
		defer cancel()
		if updateErr != nil {
			log.Panic(updateErr)
		}

		//updating trade history
		var newTrade *models.Orders = new(models.Orders)

		newTrade.OrderID = primitive.NewObjectID()
		newTrade.UserID = userID
		newTrade.Status = "EXECUTED"
		newTrade.Pair = "BTC/USD"
		newTrade.Type = "MARKET"
		newTrade.Side = "BUY"
		newTrade.Price = marketPrice
		newTrade.Amount = totalUSD.TotalUSD / newTrade.Price
		newTrade.Total = totalUSD.TotalUSD
		newTrade.OpenedAt = time.Now().Unix()
		newTrade.ExecutedAt = time.Now().Unix()

		_, insertErr := ordersCollection.InsertOne(ctx, newTrade)
		defer cancel()
		if insertErr != nil {
			log.Panic(insertErr)
		}

		c.JSON(http.StatusOK, gin.H{"success": fmt.Sprint("BTC bought at ", marketPrice)})
	}
}

func MarketSell() gin.HandlerFunc {
	return func(c *gin.Context) {
		type Market struct {
			BTCamount float32 `json:"amount"`
		}

		var btcAmount *Market = new(Market)

		//bind json from request
		if bindErr := c.ShouldBindJSON(&btcAmount); bindErr != nil {
			log.Panic(bindErr)
		}

		//getting user id from request context
		id := c.GetString("id")
		userID, idErr := primitive.ObjectIDFromHex(id)
		if idErr != nil {
			log.Panic(idErr)
		}

		//checking user's balance form db collection
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		var dbBalance *models.Users = new(models.Users)
		opts := options.FindOne().SetProjection(bson.M{"fiat": 1, "bitcoin": 1})
		dbErr := userCollection.FindOne(ctx, bson.M{"_id": userID}, opts).Decode(&dbBalance)
		defer cancel()
		if dbErr != nil {
			log.Panic(dbErr)
		}

		if dbBalance.Bitcoin < 0.0004 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient BTC"})
			return
		}

		if dbBalance.Bitcoin < btcAmount.BTCamount {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient BTC"})
			return
		}

		//getting current market price with binance api
		marketPrice, priceErr := helper.CurrentBTCPrice()
		if priceErr != nil {
			log.Panic(priceErr)
		}

		//updating balance in user collection
		_, updateErr := userCollection.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$set": bson.M{
			"fiat":    dbBalance.Fiat + (btcAmount.BTCamount * marketPrice),
			"bitcoin": dbBalance.Bitcoin - btcAmount.BTCamount,
		}})
		defer cancel()
		if updateErr != nil {
			log.Panic(updateErr)
		}

		//updating trade history
		var newTrade *models.Orders = new(models.Orders)

		newTrade.OrderID = primitive.NewObjectID()
		newTrade.UserID = userID
		newTrade.Status = "EXECUTED"
		newTrade.Pair = "BTC/USD"
		newTrade.Type = "MARKET"
		newTrade.Side = "SELL"
		newTrade.Price = marketPrice
		newTrade.Amount = btcAmount.BTCamount
		newTrade.Total = marketPrice * btcAmount.BTCamount
		newTrade.OpenedAt = time.Now().Unix()
		newTrade.ExecutedAt = time.Now().Unix()

		_, insertErr := ordersCollection.InsertOne(ctx, newTrade)
		defer cancel()
		if insertErr != nil {
			log.Panic(insertErr)
		}
		//response
		c.JSON(http.StatusOK, gin.H{"success": fmt.Sprint("BTC sold at ", marketPrice)})
	}
}

func CancelOrderById() gin.HandlerFunc {
	return func(c *gin.Context) {
		type CancelID struct {
			OrderID string `json:"order_id"`
		}

		//bind json from request
		var Cancel *CancelID = new(CancelID)
		if bindErr := c.ShouldBindJSON(&Cancel); bindErr != nil {
			log.Panic(bindErr)
		}

		if Cancel.OrderID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Order ID"})
		}
		//string to objectID
		orderID, idErr := primitive.ObjectIDFromHex(Cancel.OrderID)
		if idErr != nil {
			log.Panic(idErr)
		}

		//getting details from open order collection
		var dbDetails *models.Orders = new(models.Orders)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		dbErr := ordersCollection.FindOne(ctx, bson.M{"_id": orderID, "status": "OPEN"}).Decode(&dbDetails)
		defer cancel()
		if dbErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": dbErr.Error()})
			log.Panic(dbErr)
			return
		}

		//getting user id from request context
		id := c.GetString("id")
		userID, idErr := primitive.ObjectIDFromHex(id)
		if idErr != nil {
			log.Panic(idErr)
		}

		//updating balance in user collection
		if dbDetails.Side == "BUY" {
			_, updateErr := userCollection.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$inc": bson.M{
				"fiat": dbDetails.Total,
			}})
			defer cancel()
			if updateErr != nil {
				log.Panic(updateErr)
			}
		}

		if dbDetails.Side == "SELL" {
			_, updateErr := userCollection.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$inc": bson.M{
				"bitcoin": dbDetails.Amount,
			}})
			defer cancel()
			if updateErr != nil {
				log.Panic(updateErr)
			}
		}

		//deleting order in orders collection
		_, deErr := ordersCollection.DeleteOne(ctx, bson.M{"_id": orderID})
		if dbErr != nil {
			log.Panic(deErr)
		}

		c.JSON(http.StatusOK, gin.H{"success": fmt.Sprint("cancelled order id ", orderID.Hex())})
	}
}
