package types

import (
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JWTUser struct {
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}

func (j *JWTUser) MongoUserId() (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(j.UserId)
}
