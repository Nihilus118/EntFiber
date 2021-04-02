package main

import (
	"EntFiber/db"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Connect to Database using PGX
	db.Init()
	defer db.DB.Close()
	// Migrate the Schema
	if err := db.DB.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Create the Fiberapp
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "EST",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		ctx := context.Background()
		users, err := db.DB.User.Query().All(ctx)
		if err != nil {
			log.Fatal(err)
		}
		return c.JSON(users)
	})

	app.Listen(":3000")
}
