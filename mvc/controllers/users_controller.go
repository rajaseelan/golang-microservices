package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/rajaseelan/golang-microservices/mvc/services"
	"github.com/rajaseelan/golang-microservices/mvc/utils"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {

	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "userId must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(jsonValue))
		return
	}

	user, apiErr := services.GetUser(userId)
	if err != nil {
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(jsonValue))
		return
	}

	// return user to client
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
