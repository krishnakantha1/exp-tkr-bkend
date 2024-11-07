package mongodb

import (
	"context"
	"log"

	"github.com/krishnakantha1/expenseTrackerBackend/assert"
	"github.com/krishnakantha1/expenseTrackerBackend/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client            *mongo.Client
	databaseName      string
	expenseCollection string
	userCollection    string
}

func (mongodb *MongoDB) createIndexes() {
	var indexModel mongo.IndexModel
	var name string
	var err error

	//index for expense collection
	//unique index over user_id, year, month
	expense := mongodb.client.Database(mongodb.databaseName).Collection(mongodb.expenseCollection)
	indexModel = mongo.IndexModel{
		Keys: bson.D{
			{Key: "user_id", Value: 1},
			{Key: "year", Value: 1},
			{Key: "month", Value: 1},
		},
		Options: options.Index().SetUnique(true),
	}
	name, err = expense.Indexes().CreateOne(context.Background(), indexModel)
	assert.Error(err, "Issue while creating index on expense collection over user_id, year, month")
	log.Println("Created/found index on expense collection", name)

	//index for user collection
	//unique index over email
	user := mongodb.client.Database(mongodb.databaseName).Collection(mongodb.userCollection)
	indexModel = mongo.IndexModel{
		Keys: bson.D{
			{Key: "email", Value: 1},
		},
		Options: options.Index().SetUnique(true),
	}
	name, err = user.Indexes().CreateOne(context.Background(), indexModel)
	assert.Error(err, "Issue while creating index on user collection over email")
	log.Println("Created/found index on user collection", name)
}

func Init(envURL string, envDataBaseName string) *MongoDB {

	url, err := utils.GetEnv(envURL)
	assert.Error(err, "Issue")

	databaseName, err := utils.GetEnv(envDataBaseName)
	assert.Error(err, "Issue")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	assert.Error(err, "Problem while creating client for MongoDB")

	mongodb := &MongoDB{
		client:            client,
		databaseName:      databaseName,
		expenseCollection: "user_expense",
		userCollection:    "user_details",
	}

	mongodb.createIndexes()

	return mongodb
}
