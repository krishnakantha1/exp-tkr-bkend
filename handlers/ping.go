package handlers

import (
	"net/http"

	da "github.com/krishnakantha1/expenseTrackerBackend/dataaccess"
	"github.com/krishnakantha1/expenseTrackerBackend/types"
	"github.com/krishnakantha1/expenseTrackerBackend/utils"
)

func Ping(_ da.DataAccess, w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	response := types.ApiPingResonse{}
	response.WithSuccess(id)
	utils.SuccessResponse(w, response)
}
