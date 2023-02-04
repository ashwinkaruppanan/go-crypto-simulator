package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"ashwin.com/go-crypto-simulator/database"
	"ashwin.com/go-crypto-simulator/models"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (hashedPassword string) {
	passwordHash, hashErr := bcrypt.GenerateFromPassword([]byte(password), 14)
	if hashErr != nil {
		log.Panic(hashErr)
	}
	return string(passwordHash)
}

func VerifyPassword(loginPassword string, dbPassword string) bool {
	if passErr := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(loginPassword)); passErr != nil {
		return false
	}
	return true
}

type TokenClaims struct {
	UID  string
	Name string
	jwt.StandardClaims
}

var secretString string = os.Getenv("SECRET_KEY")

func GenerateToken(name string, userID string) (token string) {
	claims := &TokenClaims{
		UID:  userID,
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	token, tokenErr := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretString))
	if tokenErr != nil {
		log.Panic(tokenErr)
	}
	return token
}

func ValidateToken(token string) (claims *TokenClaims, msg string) {
	tkn, tknErr := jwt.ParseWithClaims(token, &TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretString), nil
	})

	if tknErr != nil {
		log.Panic(tknErr)
	}

	claims, ok := tkn.Claims.(*TokenClaims)
	if !ok {
		msg = "Invlaid Token"
		return
	}

	if claims.ExpiresAt < time.Now().Unix() {
		msg = "Token Expird"
		return
	}

	return claims, msg
}

func CurrentBTCPrice() (float32, error) {
	type BinanceAPI struct {
		Symbol string `json:"symbol"`
		Price  string `json:"price"`
	}
	response, httpErr := http.Get("https://api.binance.com/api/v3/ticker/price?symbol=BTCUSDT")
	if httpErr != nil {
		return 0, httpErr
	}
	defer response.Body.Close()

	body, bodyErr := io.ReadAll(response.Body)
	if bodyErr != nil {
		log.Panic(bodyErr)
	}

	data := &BinanceAPI{}
	if jsonErr := json.Unmarshal(body, &data); jsonErr != nil {
		log.Panic(jsonErr)
	}

	marketPrice, convErr := strconv.ParseFloat(data.Price, 32)
	if convErr != nil {
		log.Panic(convErr)
	}

	return float32(marketPrice), nil
}

func Timer() {
	next := time.Now().Add(30 * time.Second)
	for {
		fmt.Println("Waiting for the next execution...")
		now := time.Now()
		if now.After(next) {
			ExecuteOpenOrder()
			next = next.Add(30 * time.Second)
		}
		time.Sleep(next.Sub(now))
	}
}

var ordersCollection *mongo.Collection = database.OpenCollection(database.Client, "orders")
var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")

func ExecuteOpenOrder() {
	//getting current BTC price with binance API
	currentPrice, err := CurrentBTCPrice()
	if err != nil {
		log.Panic(err)
	}

	//getting open orders details
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	buyFilter := bson.M{"price": bson.M{"$gt": currentPrice, "$lt": 0.01 * currentPrice}, "status": "OPEN", "side": "BUY"}
	cursor, dbErr := ordersCollection.Find(ctx, buyFilter)
	defer cancel()
	if dbErr != nil {
		log.Panic(dbErr)
	}

	var dbBuyResult []models.Orders
	if cursorErr := cursor.All(ctx, &dbBuyResult); cursorErr != nil {
		log.Panic(cursorErr)
	}

	sellFilter := bson.M{"price": bson.M{"$lt": currentPrice, "$gt": 0.01 * currentPrice}, "status": "OPEN", "side": "SELL"}
	cursor, dbErr = ordersCollection.Find(ctx, sellFilter)
	defer cancel()
	if dbErr != nil {
		log.Panic(dbErr)
	}

	var dbSellResult []models.Orders
	if cursorErr := cursor.All(ctx, &dbSellResult); cursorErr != nil {
		log.Panic(cursorErr)
	}

	// if no orders to execute return from the function
	if len(dbBuyResult) == 0 && len(dbSellResult) == 0 {
		log.Println("No trade to execute...")
		return
	}
	// updating user balance in user collection
	//updating buy orders

	for _, val := range dbBuyResult {
		_, dbErr := userCollection.UpdateOne(ctx, bson.M{"_id": val.UserID}, bson.M{
			"$inc": bson.M{"bitcoin": val.Amount, "fiat": -val.Total},
		})
		defer cancel()
		if dbErr != nil {
			log.Panic(dbErr)
		}

		log.Println(val.OrderID.Hex(), "executed BTC price at", val.Price, "for", val.Total, "USD")

	}

	//updating sell orders

	for _, val := range dbSellResult {
		_, dbErr := userCollection.UpdateOne(ctx, bson.M{"_id": val.UserID}, bson.M{
			"$inc": bson.M{"bitcoin": -val.Amount, "fiat": val.Total},
		})
		defer cancel()
		if dbErr != nil {
			log.Panic(dbErr)
		}

		log.Println(val.OrderID.Hex(), "executed BTC price at", val.Price, "for", val.Total, "USD")

	}

	//updating executed orders in order collection
	if len(buyFilter) != 0 {
		_, upErr := ordersCollection.UpdateMany(ctx, buyFilter, bson.M{"$set": bson.M{
			"status":      "EXECUTED",
			"executed_at": time.Now().Unix(),
		}})
		defer cancel()
		if upErr != nil {
			log.Panic(upErr)
		}
	}

	if len(sellFilter) != 0 {
		_, upErr := ordersCollection.UpdateMany(ctx, sellFilter, bson.M{"$set": bson.M{
			"status":      "EXECUTED",
			"executed_at": time.Now().Unix(),
		}})
		defer cancel()
		if upErr != nil {
			log.Panic(upErr)
		}
	}
}
