package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/krishnakantha1/expenseTrackerBackend/dataaccess"
	"github.com/krishnakantha1/expenseTrackerBackend/types"
	"github.com/krishnakantha1/expenseTrackerBackend/utils"
)

/*
Function to handle ingestion of Expense data from clients.

params:

	d dataaccess.DataAccessInterface
	w http.ResponseWriter
	r *http.Request
*/
func ExpenseIngestion(d dataaccess.DataAccess, w http.ResponseWriter, r *http.Request) {
	reqBody := types.ApiExpenseIngestionRequest{}
	response := types.ApiExpenseIngestionResonse{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqBody)

	if err != nil {
		utils.BadRequestResponse(w, "Not able to parse request body")
		return
	}

	jwtStr, err := utils.GetJWTFromHeader(r)
	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	jwtUser, err := utils.GetUserFromJWT(jwtStr)
	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	saveCount, err := d.PersistExpenses(jwtUser, reqBody.Expenses)
	if err != nil {
		response.WithError(fmt.Sprintf("Error in saving expenses. %v", err))
		utils.SuccessResponse(w, response)
	} else {
		response.WithSuccess(saveCount)
		utils.SuccessResponse(w, response)
	}
}
