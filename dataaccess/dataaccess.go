package dataaccess

import "github.com/krishnakantha1/expenseTrackerBackend/types"

type DataAccess interface {
	GetUserEmail(email string) *types.User
	SaveUser(user *types.User) (*types.User, error)
	PersistExpenses(*types.JWTUser, []*types.ApiExpenseMessage) (int, error)
}
