package routes

import (
	"myapp/controllers"

	"github.com/gofiber/fiber/v2"
	//"github.com/gofiber/websocket"
	"github.com/gofiber/websocket/v2"
)

// SetupRoutes defines all the routes for the application.
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/items", controllers.GetItems)          // Get all items
	api.Post("/items", controllers.CreateItem)       // Create a new item
	api.Put("/items/:id", controllers.UpdateItem)    // Update an existing item
	api.Delete("/items/:id", controllers.DeleteItem) // Delete an item

	api.Get("/user", controllers.GetTbluser)           // Get all items
	api.Post("/user", controllers.CreateTbluser)       // Create a new item
	api.Put("/user/:id", controllers.UpdateTbluser)    // Update an existing item
	api.Delete("/user/:id", controllers.DeleteTbluser) // Delete an item

	api.Post("/userLogin", controllers.ValidateUser) // Get Login

	api.Get("/counter", controllers.GetTblcounter)           // Get all items
	api.Post("/counter", controllers.CreateTblcounter)       // Create a new item
	api.Put("/counter/:id", controllers.UpdateTblcounter)    // Update an existing item
	api.Delete("/counter/:id", controllers.DeleteTblcounter) // Delete an item

	api.Get("/department", controllers.GetTbldepartment)           // Get all items
	api.Post("/department", controllers.CreateTbldepartment)       // Create a new item
	api.Put("/department/:id", controllers.UpdateTbldepartment)    // Update an existing item
	api.Delete("/department/:id", controllers.DeleteTbldepartment) // Delete an item

	api.Put("/department/increment/:id", controllers.IncrementTbldepartmentNumber)

	api.Get("/que", controllers.GetTblque)           // Get all items
	api.Post("/que", controllers.CreateTblque)       // Create a new item
	api.Put("/que/:id", controllers.UpdateTblque)    // Update an existing item
	api.Delete("/que/:id", controllers.DeleteTblque) // Delete an item

	app.Get("/video1", controllers.VideoHandler1)
	app.Get("/video2", controllers.VideoHandler2)

	websocketController := controllers.NewWebSocketController()
	// Define the WebSocket route
	app.Get("/ws", websocket.New(websocketController.HandleWebSocketConnection))
}
