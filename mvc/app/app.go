package app

import(
"net/http"
"github.com/alishaagupta/microservices-with-go/mvc/controllers"
"github.com/gin-gonic/gin"

)

var(
	router *gin.Engine
)

func init(){

	router = gin.Default()
	// router = gin.New()
}

func StartApp(){

 mapUrls()

 // Launch the server
 if err := router.Run(":7000"); err != nil {
	 panic(err)
 } 

}