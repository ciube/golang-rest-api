package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// creo applicazione
	app := fiber.New()
	// faccio usare ad app un Logger
	app.Use(logger.New())

	// metodo GET che risponde al path "/"
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the endpoint 😉",
		})
	})

	// app sente su localhost:8080 e se ci sta un errore popola err
	err := app.Listen("localhost:8080")

	// se err è diverso da null stoppo esecuzione della goroutine con panic
	if err != nil {
		panic(err)
	}
}
