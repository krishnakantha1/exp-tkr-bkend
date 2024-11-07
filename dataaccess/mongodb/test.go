package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

type MyObj struct {
	List_id     string    `bson:"list_id"`
	LastUpdated time.Time `bson:"last_upadted"`
	Value       int       `bson:"Value"`
}

type Test struct {
	Year  int     `bson:"year"`
	Month int     `bson:"month"`
	List  []MyObj `bson:"list"`
}

func (m *MongoDB) TestInsert() {
	lis := make([]MyObj, 0)
	lis = append(lis, MyObj{
		List_id:     "test1",
		LastUpdated: time.Now(),
		Value:       5,
	},
		MyObj{
			List_id:     "test2",
			LastUpdated: time.Now(),
			Value:       18,
		},
		MyObj{
			List_id:     "test3",
			LastUpdated: time.Now(),
			Value:       11,
		},
	)

	data := Test{
		Year:  1,
		Month: 0,
		List:  lis,
	}

	coll := m.client.Database("test_expenseTracker").Collection("test_1")

	Values := []interface{}{
		data,
	}

	coll.InsertMany(context.TODO(), Values)
}

func (m *MongoDB) TestUpsert() {
	coll := m.client.Database(m.databaseName).Collection("test_1")

	filter := bson.D{{Key: "year", Value: 1}, {Key: "month", Value: 0}, {Key: "list.list_id", Value: bson.D{{Key: "$ne", Value: "test5"}}}}

	value := bson.D{{
		Key: "$addToSet",
		Value: bson.D{{
			Key: "list",
			Value: MyObj{
				List_id:     "test5",
				LastUpdated: time.Now(),
				Value:       10,
			}}},
	}}

	res, err := coll.UpdateMany(context.TODO(), filter, value)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.ModifiedCount)
	}
}
