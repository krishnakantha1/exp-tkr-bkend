package types

import (
	"time"

	"github.com/krishnakantha1/expenseTrackerBackend/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        string
	Username  string
	Email     string
	Password  string
	CreatedOn time.Time
	UpdatedOn time.Time
	AesTest   string
}

func (u *User) ToMongoUser() *MongoUser {
	id, err := primitive.ObjectIDFromHex(u.ID)
	assert.Info(err, "Couldn't convert MongoUser to User")

	return &MongoUser{
		ID:        id,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		CreatedOn: u.CreatedOn,
		UpdatedOn: u.UpdatedOn,
		AesTest:   u.AesTest,
	}
}

func (u *User) ToJWTUser() *JWTUser {
	return &JWTUser{
		UserId: u.ID,
	}
}
