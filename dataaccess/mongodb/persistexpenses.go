package mongodb

import (
	"context"

	"github.com/krishnakantha1/expenseTrackerBackend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
structure to group types.ApiExpenseMessage sent by client.
the item in these arrays might have entries belonging to different year/ months.
they need to be grouped based on year and month.
*/
type groupedExpenses struct {
	year     int
	month    int
	expenses []*types.MongoExpense
}

/*
DataAccessInterface implementation of PersistExpenses.

params:

	user *types.JWTUser
	expenseMessages []*types.ApiExpenseMessage

returns:

	int
	error
*/
func (m *MongoDB) PersistExpenses(user *types.JWTUser, expenseMessages []*types.ApiExpenseMessage) (int, error) {
	userID, err := user.MongoUserId()
	if err != nil {
		return -1, err
	}

	finalCount := 0

	for _, item := range groupExpensesOnMonth(expenseMessages) {
		count, _ := m.upsertExpenses(userID, item.year, item.month, item.expenses)
		finalCount += count
	}

	return finalCount, nil
}

/*
given a slice of types.ApiExpenseMessage returns a slice of grouped expense

params:

	expenses []*types.ApiExpenseMessage

return:

	[]*groupedExpenses
*/
func groupExpensesOnMonth(expenses []*types.ApiExpenseMessage) []*groupedExpenses {
	res := make([]*groupedExpenses, 0)
	mapper := make(map[int]*groupedExpenses)

	for _, expense := range expenses {
		mongoExpense := expense.ExpsenseEntry.ToMongoExpense()

		year := mongoExpense.ExpenseDate.Year()
		month := int(mongoExpense.ExpenseDate.Month())

		key := (month)*10000 + year

		if _, ok := mapper[key]; !ok {
			groupedEntry := &groupedExpenses{
				year:     year,
				month:    month,
				expenses: make([]*types.MongoExpense, 0, 1),
			}

			res = append(res, groupedEntry)
			mapper[key] = groupedEntry
		}

		groupedEntry := mapper[key]
		groupedEntry.expenses = append(groupedEntry.expenses, mongoExpense)
	}

	return res
}

/*
upserts the given expenses into mongo db expense collection.
the expense data for all expenses in the slice should be of the same month and year.

params:

	userID primitive.ObjectID
	year int
	month int
	expenses []*types.MongoExpense

returns:

	int
	error
*/
func (m *MongoDB) upsertExpenses(userID primitive.ObjectID, year int, month int, expenses []*types.MongoExpense) (int, error) {
	coll := m.client.Database(m.databaseName).Collection(m.expenseCollection)

	models := []mongo.WriteModel{}

	// most likely the below will fail as there can only be one Expense Document for a userID, year, month.
	// we want to make sure it exists to push in the expenses into it
	models = append(models, createExpenseDocument(userID, year, month))

	for _, expense := range expenses {
		models = append(models, createExpenseEntry(userID, year, month, expense))
	}

	options := options.BulkWrite().SetOrdered(true)

	result, err := coll.BulkWrite(context.Background(), models, options)

	return int(result.InsertedCount), err
}

/*
creates an Insert One Model (mongo.WriteModel) for the Expense document
params:

	userID primitive.ObjectID
	year int
	month int

returns:

	mongo.WriteModel
*/
func createExpenseDocument(userID primitive.ObjectID, year int, month int) mongo.WriteModel {
	return mongo.NewInsertOneModel().SetDocument(
		types.MongoExpenseDocument{
			UserID:   userID,
			Year:     year,
			Month:    month,
			Expenses: make([]*types.MongoExpense, 0),
		},
	)

}

/*
creates a Update One Model (mongo.WriteModel) for Expense Entries
params:

	userID primitive.ObjectID
	year int
	month int
	expense *types.MongoExpense

return:

	mongo.WriteModel
*/
func createExpenseEntry(userID primitive.ObjectID, year int, month int, expense *types.MongoExpense) mongo.WriteModel {
	filter := bson.M{
		"user_id": userID,
		"year":    year,
		"month":   month,
		"expenses.uri": bson.M{
			"$ne": expense.URI,
		},
	}

	update := bson.M{
		"$push": bson.M{
			"expenses": expense,
		},
	}

	return mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update)
}
