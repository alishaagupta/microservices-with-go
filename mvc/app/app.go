package app

import(
"net/http"
"github.com/alishaagupta/microservices-with-go/mvc/controllers"

)


func StartApp(){

 http.HandleFunc("/users", controllers.GetUser) ;

 // Launch the server
 if err := http.ListenAndServe(":7000", nil); err != nil {
	 panic(err)
 } 

}