package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/krishnakantha1/expenseTrackerBackend/dataaccess"
	"github.com/krishnakantha1/expenseTrackerBackend/types"
	"github.com/krishnakantha1/expenseTrackerBackend/utils"
)

func Login(d dataaccess.DataAccess, w http.ResponseWriter, r *http.Request) {
	reqBody := types.ApiLoginCredentialsRequest{}
	response := types.ApiLoginCredentialsResponse{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqBody)

	if err != nil {
		utils.BadRequestResponse(w, "not able to parse request body")
		return
	}

	user := d.GetUserEmail(reqBody.Email)
	if user == nil {
		response.WithError("username or password is wrong")
		utils.SuccessResponse(w, response)
		return
	}

	jwt, err := utils.GetJWTFromUser(user.ToJWTUser())
	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	response.WithSuccess(jwt)
	utils.SuccessResponse(w, response)
}

func LoginWithJWT(d dataaccess.DataAccess, w http.ResponseWriter, r *http.Request) {
	jwt, err := utils.GetJWTFromHeader(r)
	response := types.ApiLoginCredentialsResponse{}

	if err != nil {
		response.WithError("could not obtain JWT from bearer")
		utils.SuccessResponse(w, response)
		return
	}

	user, err := utils.GetUserFromJWT(jwt)

	if err != nil {
		response.WithError("bad jwt")
		utils.SuccessResponse(w, response)
		return
	}

	jwt, err = utils.GetJWTFromUser(user)

	if err != nil {
		utils.BadRequestResponse(w, err.Error())
		return
	}

	response.WithSuccess(jwt)
	utils.SuccessResponse(w, response)
}

func Register(d dataaccess.DataAccess, w http.ResponseWriter, r *http.Request) {

}
