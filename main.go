package main

import (
	"log"

	"github.com/dortaedward/financeTracker/types"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/django/v3"
)

const(
	addr = ":6969"
)

func main(){

	engine := django.New("./views", ".html")
	engine.Reload(true)
	app := fiber.New(fiber.Config{
		Views: engine,
		PassLocalsToViews: true,
	})
	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(logger.New())

	app.Static("/static","./static/")

	api := types.CreateServer(addr,app)
	api.Run()

	err := app.Listen(addr); if err != nil {
		log.Fatal("ERROR: There was an error creating the server \n,",err)
	}
}
