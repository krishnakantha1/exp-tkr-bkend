package types

import (
	"time"
)

/*
request body for expense ingestion api request
*/
type ApiExpenseIngestionRequest struct {
	Count    int                  `json:"count"`
	Expenses []*ApiExpenseMessage `json:"expenses"`
}

type ApiExpenseMessage struct {
	ExpsenseEntry *ApiExpsenseEntry `json:"expense_entry"`
	RawMessage    *ApiRawMessage    `json:"raw_message"`
}

/*
details for individual expense sent by client
*/
type ApiExpsenseEntry struct {
	URI             string `json:"uri"`
	Bank            string `json:"bank"`
	EncryptedAmount string `json:"encrypted_amount"`
	ExpenseDate     int64  `json:"expensedate_long"`
	ExpenseType     string `json:"expense_type"`
	ExpenseTag      string `json:"tag"`
}

func (aee *ApiExpsenseEntry) ToMongoExpense() *MongoExpense {
	return &MongoExpense{
		URI:             aee.URI,
		Bank:            aee.Bank,
		AmountEncrypted: aee.EncryptedAmount,
		ExpenseDate:     time.Unix(aee.ExpenseDate/1000, (aee.ExpenseDate%1000)*1000_000),
		IngestedOn:      time.Now(),
		ExpenseType:     aee.ExpenseType,
		ExpenseTag:      aee.ExpenseTag,
	}
}

/*
raw message of expense (SMS) sent from client
*/
type ApiRawMessage struct {
	Raw string `json:"raw"`
}

/*
request body for login with credentials api request
*/
type ApiLoginCredentialsRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

/*
fields for all api responses.
*/
type ApiResonse struct {
	Error        bool   `json:"error"`
	ErrorMessage string `json:"error_message"`
}

func (r *ApiResonse) WithError(msg string) {
	r.Error = true
	r.ErrorMessage = msg
}

/*
response message for ping
*/
type ApiPingResonse struct {
	ApiResonse
	ID string `json:"id"`
}

func (r *ApiPingResonse) WithSuccess(id string) {
	r.ID = id
}

/*
response message for expense ingestion api request
*/
type ApiExpenseIngestionResonse struct {
	ApiResonse
	Count int `json:"ingestion_count"`
}

func (r *ApiExpenseIngestionResonse) WithSuccess(count int) {
	r.Count = count
}

/*
response message for login via credentials api request
*/
type ApiLoginCredentialsResponse struct {
	ApiResonse
	JWT string `json:"jwt"`
}

func (r *ApiLoginCredentialsResponse) WithSuccess(jwt string) {
	r.JWT = jwt
}
