package main

import (
	"log"
	"myapp/models"
	"myapp/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberSwagger "github.com/swaggo/fiber-swagger"

	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	//"gorm.io/gorm/logger"
)

func main() {
	// Create a new fiber instance
	app := fiber.New()

	// Serve Swagger UI
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	app.Use(logger.New())

	// Setup CORS middleware
	app.Use(cors.New())

	// Setup database connection
	//db, err := gorm.Open("sqlite3", "./gorm.db")

	dsn := "sqlserver://test:test2019@localhost:1434?database=GOOPENAPI_UAT"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect to the database")
	}
	models.InitializeDatabase(db)

	// Pass the database instance to the app's context
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	// Setup routes
	routes.SetupRoutes(app)

	// swag.Init()

	// Start server
	log.Fatal(app.Listen(":3000"))
}
