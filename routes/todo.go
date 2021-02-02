package routes

import (
	"github.com/ciube/golang-rest-api/controllers"
	"github.com/gofiber/fiber/v2"
)

// TodoRoute route ai metodi Todo
func TodoRoute(route fiber.Router) {
	// metodo GET che mappa il metodo GetTodos
	route.Get("", controllers.GetTodos)
}
