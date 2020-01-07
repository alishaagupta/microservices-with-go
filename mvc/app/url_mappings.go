package app

import(
	"github.com/alishaagupta/microservices-with-go/mvc/controllers"

	"net/http"
)

func mapUrls(){

	// http.HandleFunc("/users", controllers.GetUser) ;

	router.GET("/users/:user_id", controllers.GetUser)
}