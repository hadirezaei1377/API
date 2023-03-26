package main

/*

import (
	"API/config"
	"API/routing"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

// related to middlewares
func RootLevel(next echo.HandlerFunc) echo.HandlerFunc {
	fmt.Println("RootLevel")
	return next
}

func AfterRouter(next echo.HandlerFunc) echo.HandlerFunc {
	fmt.Println("AfterRouter")
	return next
}


// next is a paramter that will be passed to next middlewares

// The initial architecture of the project

func main() {

	// get config , Recall information about the config
	err := config.GetConfig()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Server port : ", config.AppConfig.Server.Port)

	// init server or create it
	server := echo.New()
	// Setting the validator on the server
    server.Validator = &Utility.CustomValidator{Validator: validator.New()}
	// server jobs

	   // know your routes
	   // set  middleware
	     // creating middleware
		 server.Pre(RootLevel)
	     server.Use(AfterRouter)



	// routing
	routing.SetRouting(server)

	// start server or implement it
	server.Start(":" + config.AppConfig.Server.Port)
}

*/
