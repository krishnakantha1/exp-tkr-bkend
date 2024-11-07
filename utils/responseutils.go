package utils

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/krishnakantha1/expenseTrackerBackend/assert"
	"github.com/krishnakantha1/expenseTrackerBackend/types"
)

func BadRequestResponse(w http.ResponseWriter, msg ...string) {
	writeErrorResponse(http.StatusBadRequest, w, msg...)
}

func ServerErrorResponse(w http.ResponseWriter, msg ...string) {
	writeErrorResponse(http.StatusInternalServerError, w, msg...)
}

func SuccessResponse(w http.ResponseWriter, msgStruct any) error {
	msg, err := json.Marshal(msgStruct)
	assert.Error(err, "Unexpected error while creating SuccessResponse message")

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(msg)

	return err
}

func combineString(strs ...string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	var sb strings.Builder

	for i := 0; i < len(strs); i++ {
		sb.WriteString(strs[i])

		if i != len(strs)-1 {
			sb.WriteByte('\n')
		}
	}

	return sb.String()
}

func writeErrorResponse(code int, w http.ResponseWriter, msg ...string) {
	response := types.ApiResonse{}
	response.WithError(combineString(msg...))

	rb, err := json.Marshal(response)
	assert.Error(err, "Unexpected error while creating BadRequestResponse message")

	http.Error(w, string(rb), code)
}
