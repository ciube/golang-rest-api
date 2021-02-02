package controllers

import (
	"github.com/gofiber/fiber/v2"
)

// Todo type
// definiamo il tipo Todo
type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// inizializziamo qualche dato
var todos = []*Todo{
	{
		ID:        1,
		Title:     "Walk the dog 🦮",
		Completed: false,
	},
	{
		ID:        2,
		Title:     "Walk the cat 🐈",
		Completed: false,
	},
}

// GetTodos Metodo GET per recuperare tutti i Todos
func GetTodos(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todos": todos,
		},
	})
}
