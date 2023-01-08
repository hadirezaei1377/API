package main

import (
	"API/config"
	"API/routing"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

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

	// server jobs
	/*
	   // know your routes
	   // set  middleware

	*/

	// routing
	routing.SetRouting(server)

	// start server or implement it
	server.Start(":" + config.AppConfig.Server.Port)
}
