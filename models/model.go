package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID        primitive.ObjectID `bson:"_id"`
	FullName  string             `json:"full_name" bson:"full_name" validate:"required,min=5,max=25"`
	Email     string             `json:"email" bson:"email" validate:"email,required"`
	Password  string             `json:"password" bson:"password" validate:"required"`
	Fiat      float64            `json:"fiat" bson:"fiat"`
	Bitcoin   float64            `json:"bitcoin" bson:"bitcoin"`
	JoinedAt  int64              `json:"joined_at" bson:"joined_at"`
	UpdatedAt int64              `json:"updated_at" bson:"updated_at"`
}

type OpenOrders struct {
	ID     primitive.ObjectID `bson:"_id"`
	Date   int64              `json:"date" bson:"date"`
	Pair   string             `json:"pair" bson:"pair"`
	Type   string             `json:"type" bson:"type"`
	Side   string             `json:"side" bson:"side"`
	Price  float64            `json:"price" bson:"price"`
	Amount float64            `json:"amount" bson:"amount"`
	Total  float64            `json:"total" bson:"total"`
}

type Transaction struct {
	ID     primitive.ObjectID `bson:"_id"`
	Date   int64              `json:"date" bson:"date"`
	Pair   string             `json:"pair" bson:"pair"`
	Type   string             `json:"type" bson:"type"`
	Side   string             `json:"side" bson:"side"`
	Price  float64            `json:"price" bson:"price"`
	Amount float64            `json:"amount" bson:"amount"`
	Total  float64            `json:"total" bson:"total"`
	Fee    float64            `json:"fee" bson:"fee"`
}
