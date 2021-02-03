package controllers

import (
	"fmt"
	"strconv"

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

// CreateTodo Metodo POST per creare un nuovo Todo
func CreateTodo(c *fiber.Ctx) error {
	// Request type
	//tipo che definisce i campi della richiesta
	type Request struct {
		Title string `json:"title"`
	}

	// inizializzo il body della richiesta
	var body Request

	// bodyParse del body
	err := c.BodyParser(&body)

	// se err è diverso da null ritorno 500
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}

	// definisco il todo da agganciare all'array dei todos
	todo := &Todo{
		ID:        len(todos) + 1,
		Title:     body.Title,
		Completed: false,
	}

	todos = append(todos, todo)

	// ritorno lo stato created e il todo appena generato
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todo": todo,
		},
	})
}

// GetTodo Metodo GET per recuperare un Todo a partire da un ID
func GetTodo(c *fiber.Ctx) error {
	// prendo l'id dai parametri della chiamata
	paramID := c.Params("id")

	// converto il parametro a int
	id, err := strconv.Atoi(paramID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse Id",
		})
	}

	for _, todo := range todos {
		if todo.ID == id {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"data": fiber.Map{
					"todo": todo,
				},
			})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": false,
		"message": "Todo not found",
	})
}

// UpdateTodo Metodo PUT per aggiornare un Todo
func UpdateTodo(c *fiber.Ctx) error {
	paramID := c.Params("id")

	id, err := strconv.Atoi(paramID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse id",
		})
	}

	type Request struct {
		Title     *string `json:"title"`
		Completed *bool   `json:"completed"`
	}

	var body Request
	err = c.BodyParser(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}

	var todo *Todo

	for _, t := range todos {
		if t.ID == id {
			todo = t
			break
		}
	}

	if todo.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Not Found",
		})
	}

	if body.Title != nil {
		todo.Title = *body.Title
	}

	if body.Completed != nil {
		todo.Completed = *body.Completed
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todo": todo,
		},
	})

}

func DeleteTodo(c *fiber.Ctx) error {
	// get param
	paramID := c.Params("id")

	// convert param string to int
	id, err := strconv.Atoi(paramID)

	// if parameter cannot parse
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse id",
		})
	}

	// find and delete todo
	for i, todo := range todos {
		if todo.ID == id {

			todos = append(todos[:i], todos[i+1:]...)

			return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
				"success": true,
				"message": "Deleted Succesfully",
			})
		}
	}

	// if todo not found
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": false,
		"message": "Todo not found",
	})
}
