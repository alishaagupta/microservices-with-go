package controllers 

import(
	"net/http"
	"log"
	"strconv"
	"github.com/alishaagupta/microservices-with-go/mvc/services"
	"github.com/alishaagupta/microservices-with-go/mvc/utils"

	"encoding/json"
	"github.com/gin-gonic/gin"



)

func GetUser(c *gin.Context){


	// userIdParam := req.URL.Query().Get("user_id");

	// Use c.Query("user_id") if sent as query parameter
	userIdParam := c.Param("user_id")

	log.Printf("About to process user_id %v ", userIdParam)

	userId, err := strconv.ParseInt(userIdParam, 10, 64) 

	if err != nil {

		apiErr := &utils.ApplicationError{
			Message: "user_id must be a number ",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}

		// Equivalent to writing reposne to the client

		c.JSON(apiErr.StatusCode, apiErr)

		// jsonValue, _ := json.Marshal(apiErr)
		// // Return the error to the client 
		// resp.WriteHeader(apiErr.StatusCode)
		// resp.Write(jsonValue)
		return
	}

	user, apiErr := services.GetUser(userId)

	if apiErr!= nil{
		// Handle the error

		c.JSON(apiErr.StatusCode, apiErr)
		// resp.WriteHeader(apiErr.StatusCode)
		// resp.Write([]byte(apiErr.Message))
		return
	}

	// Return user to the client

	c.JSON(http.StatusOK, user)
	// jsonValue, _ := json.Marshal(user)
	// resp.Write(jsonValue)
}