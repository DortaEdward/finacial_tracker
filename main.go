package main

import (
	"log"

	"github.com/dortaedward/financeTracker/types"
	"github.com/gofiber/fiber/v2"
)

const(
	addr = ":6969"
)

func main(){

	app := fiber.New()

	api := types.CreateServer(addr,app)
	api.Run()

	err := app.Listen(addr); if err != nil {
		log.Fatal("ERROR: There was an error creating the server \n,",err)
	}
}
