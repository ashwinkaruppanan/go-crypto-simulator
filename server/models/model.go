package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `json:"name" bson:"name,omitempty" validate:"required,min=5,max=25"`
	Email     string             `json:"email" bson:"email,omitempty" validate:"email,required"`
	Password  string             `json:"password" bson:"password,omitempty" validate:"required"`
	Fiat      float32            `json:"fiat,omitempty" bson:"fiat,omitempty,truncate"`
	Bitcoin   float32            `json:"bitcoin,omitempty" bson:"bitcoin,omitempty,truncate"`
	JoinedAt  int64              `json:"joined_at,omitempty" bson:"joined_at,omitempty"`
	UpdatedAt int64              `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type OpenOrders struct {
	ID     primitive.ObjectID `bson:"_id"`
	UserID primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	Date   int64              `json:"date" bson:"date,omitempty"`
	Pair   string             `json:"pair" bson:"pair,omitempty"`
	Type   string             `json:"type" bson:"type,omitempty"`
	Side   string             `json:"side" bson:"side,omitempty"`
	Price  float32            `json:"price" bson:"price,omitempty"`
	Amount float32            `json:"amount" bson:"amount,omitempty"`
	Total  float32            `json:"total" bson:"total,omitempty"`
}

type TradeHistory struct {
	ID     primitive.ObjectID `bson:"_id"`
	UserID primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	Date   int64              `json:"date" bson:"date,omitempty"`
	Pair   string             `json:"pair" bson:"pair,omitempty"`
	Type   string             `json:"type" bson:"type,omitempty"`
	Side   string             `json:"side" bson:"side,omitempty"`
	Price  float32            `json:"price" bson:"price,omitempty"`
	Amount float32            `json:"amount" bson:"amount,omitempty"`
	Total  float32            `json:"total" bson:"total,omitempty"`
	Fee    float32            `json:"fee" bson:"fee,omitempty"`
}
