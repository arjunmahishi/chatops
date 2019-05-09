package main

import (
	"log"

	"github.com/arjunmahishi/Chatops/commanders"
	"github.com/arjunmahishi/Chatops/config"
	"github.com/arjunmahishi/Chatops/messenger"
	"github.com/arjunmahishi/Chatops/routes"
	echo "github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var sender = messenger.NewMessenger()

func main() {
	err := commanders.SyncCommands(config.Config.CommandsPath)
	if err != nil {
		// sender.Send("spaces/AAAA8fjBKEQ", "<users/all> Chatops has Crashed: \n```"+err.Error()+"```")
		log.Fatalf(err.Error())
	}

	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} method=${method} uri=${uri} status=${status}\n",
	}))

	e.POST("/", routes.Chatbot)

	// sender.Send("spaces/AAAA8fjBKEQ", "Chatops restarted at "+time.Now().String())
	e.Logger.Fatal(e.Start(":1323"))
}
