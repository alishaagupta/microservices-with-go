package services

import(
	"github.com/alishaagupta/microservices-with-go/mvc/domain"
	"github.com/alishaagupta/microservices-with-go/mvc/utils"


)


func GetUser(userId int64) (*domain.User, *utils.ApplicationError){

	return domain.GetUser(userId)
}