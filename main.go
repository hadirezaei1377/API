package main

import (
	"fmt"
	"log"
	"toplearn-golang-api/config"
	"toplearn-golang-api/routing"

	"github.com/labstack/echo/v4"
)

func main() {

	// config
	err := config.GetConfig()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Server port : ", config.AppConfig.Server.Port)

	//init server
	server := echo.New()

	//routing
	routing.SetRouting(server)

	//middleware

	//start server
	server.Start(":" + config.AppConfig.Server.Port)
}
