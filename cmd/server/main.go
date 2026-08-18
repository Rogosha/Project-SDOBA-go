package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"

	"SDOBA/internal/config"
	"SDOBA/internal/database"
)

func main() {
	cfg := config.Load()

	db, err := database.NewPostgresPool(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		err := db.Ping(context.Background())

		if err != nil {
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"status":   "error",
				"database": "unavailable",
			})
		}

		return c.JSON(fiber.Map{
			"status":   "ok",
			"database": "ok",
			"service":  cfg.App.Name,
		})
	})

	log.Fatal(app.Listen(":" + cfg.App.Port))
}
