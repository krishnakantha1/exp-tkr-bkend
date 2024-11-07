package mongodb

import (
	"context"
	"fmt"

	"github.com/krishnakantha1/expenseTrackerBackend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m *MongoDB) GetUserEmail(email string) *types.User {
	coll := m.client.Database(m.databaseName).Collection(m.userCollection)

	filters := bson.D{
		{Key: "email", Value: email},
	}

	mongoUser := types.MongoUser{}
	err := coll.FindOne(context.Background(), filters).Decode(&mongoUser)

	if err != nil {
		return nil
	}

	return mongoUser.ToUser()
}

func (m *MongoDB) SaveUser(user *types.User) (*types.User, error) {
	coll := m.client.Database(m.databaseName).Collection(m.userCollection)

	result, err := coll.InsertOne(context.Background(), user.ToMongoUser())

	if err != nil {
		return nil, err
	}

	if id, ok := result.InsertedID.(primitive.ObjectID); ok {
		user.ID = id.Hex()
	} else {
		return nil, fmt.Errorf("error converting inserted id interface to mongo objectid")
	}

	return user, nil
}
