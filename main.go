package main

import (
	"database/sql"
	"log"

	"github.com/dortaedward/financeTracker/api"
	"github.com/dortaedward/financeTracker/types"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const(
	addr = ":6969"
)

var db *sql.DB

func main(){

	app := fiber.New(fiber.Config{
		PassLocalsToViews: true,
	})

	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(logger.New())

  db = types.ConnectDb() 

	api := api.CreateServer(addr,app, db)
	api.Run()

	err := app.Listen(addr); if err != nil {
		log.Fatal("ERROR: There was an error creating the server \n,",err)
	}
}
