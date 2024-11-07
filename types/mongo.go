package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
Details for a Mongo User document
*/
type MongoUser struct {
	ID        primitive.ObjectID `bson:"_id"`
	Username  string             `bson:"username"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	CreatedOn time.Time          `bson:"created_on"`
	UpdatedOn time.Time          `bson:"updated_on"`
	AesTest   string             `bson:"aestest"`
}

func (m *MongoUser) ToUser() *User {
	return &User{
		ID:        m.ID.Hex(),
		Username:  m.Username,
		Email:     m.Email,
		Password:  m.Password,
		CreatedOn: m.CreatedOn,
		UpdatedOn: m.UpdatedOn,
		AesTest:   m.AesTest,
	}
}

/*
Details for a Mongo Expense document
*/
type MongoExpenseDocument struct {
	ID       primitive.ObjectID `bson:"_id"`
	UserID   primitive.ObjectID `bson:"user_id"`
	Year     int                `bson:"year"`
	Month    int                `bosn:"month"`
	Expenses []*MongoExpense    `bson:"expenses"`
}

type MongoExpense struct {
	URI             string    `bson:"uri"`
	Bank            string    `bson:"bank"`
	AmountEncrypted string    `bson:"amount_encrypted"`
	ExpenseDate     time.Time `bson:"expense_date"`
	IngestedOn      time.Time `bson:"ingested_on"`
	ExpenseType     string    `bson:"expense_type"`
	ExpenseTag      string    `bson:"tag"`
}
