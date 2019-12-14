package domain

import(
	"fmt"
	"github.com/alishaagupta/microservices-with-go/mvc/utils"
	"net/http"

)

// In this example, this is the connection database with single entry only.
var (

	users = map[int64]*User{
		123: {Id: 1, FirstName: "Alisha", LastName: "Gupta", Email: "alishagupta262@gmail.com"},
	}
)


func GetUser(userId int64)(*User , *utils.ApplicationError){

	//Assume we are connecting to a database adn getting data from that

	//Here assume users is from db

	if user:= users[userId]; user != nil {

		return user, nil

	}

	// return nil, errors.New(fmt.Sprintf("user %w is not found ", userId))

	return nil, &utils.ApplicationError{
		Message: fmt.Sprintf("user %w is not found ", userId),
		StatusCode: http.StatusNotFound,
		Code: "not_found",
	}

}