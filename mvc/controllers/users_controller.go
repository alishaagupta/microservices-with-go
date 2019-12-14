package controllers ;

import(
	"net/http"
	"log"
	"strconv"
	"github.com/alishaagupta/microservices-with-go/mvc/services"
	"github.com/alishaagupta/microservices-with-go/mvc/utils"

	"encoding/json"



)

func GetUser(resp http.ResponseWriter, req *http.Request){


	userIdParam := req.URL.Query().Get("user_id");

	log.Printf("About to process user_id %v ", userIdParam)

	userId, err := strconv.ParseInt(userIdParam, 10, 64) 

	if err != nil {

		apiErr := &utils.ApplicationError{
			Message: "user_id must be a number ",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}

		jsonValue, _ := json.Marshal(apiErr)
		// Return the error to the client 
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	user, apiErr := services.GetUser(userId)

	if apiErr!= nil{
		// Handle the error

		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(apiErr.Message))
		return
	}

	// Return user to the client

	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}