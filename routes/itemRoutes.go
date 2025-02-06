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

	api.Get("/userlog", controllers.GetTbluserlog)        // Get all items
	api.Post("/userlog", controllers.CreateTbluserlog)    // Create a new item
	api.Put("/userlog/:id", controllers.UpdateTbluserlog) // Update an existing item

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
	api.Get("/quetran", controllers.GetTblquebyTran) // Get all items WAITING/CALL

	//20250204 - add queKiosk
	api.Get("/quekiosk", controllers.GetTblquebyTran) // Get all items by Transaction

	api.Get("/vid/video1", controllers.VideoHandler1)
	api.Get("/vid/video2", controllers.VideoHandler2)

	websocketController := controllers.NewWebSocketController()

	// Upload API endpoint
	api.Post("/upload", controllers.UploadFile)
	api.Post("/upload/:username", controllers.UploadFileuser)

	// Serve static files (uploaded images)
	app.Static("/uploads", "./static/uploads")

	// Download API endpoint
	api.Get("/getpic/:imagename", controllers.GetFile)
	api.Get("/download/:imagename", controllers.DownloadFile)

	// Define the WebSocket route
	app.Get("/ws", websocket.New(websocketController.HandleWebSocketConnection))

	api.Get("/download/excel", controllers.GenerateExcelFile)

}
