# C:\jsj\Exercise\8-20250112-FiberGinMVCQuasar\backend> 

go get github.com/gofiber/fiber/v2
go get github.com/jinzhu/gorm
go get github.com/jinzhu/gorm/dialects/sqlite
go get github.com/gofiber/fiber/v2/middleware/cors


backend/
  ├── controllers/
  │   └── itemController.go
  ├── models/
  │   └── itemModel.go
  ├── routes/
  │   └── itemRoutes.go
  ├── main.go

# Swagger requirement
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/fiber-swagger
go get -u github.com/gofiber/swagger

your_project/
│
├── docs/                   # Swagger generated docs
├── main.go                 # Your main API file
└── go.mod                  # Go module file

import (
	"github.com/swaggo/fiber-swagger"
	_ "your_project/docs" // This imports the generated swagger documentation
)

	// Swagger Route
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// Routes
	app.Get("/users", getUsers)
	app.Post("/users", createUser)

  // @Summary Get users
// @Description Get all users from the database
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} User
// @Router /users [get]
func getUsers(c *fiber.Ctx) error {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		return c.Status(500).SendString("Failed to get users")
	}
	return c.JSON(users)
}

// @Summary Create a new user
// @Description Add a new user to the database
// @Tags users
// @Accept json
// @Produce json
// @Param user body User true "User Data"
// @Success 201 {object} User
// @Router /users [post]
func createUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString("Invalid input")
	}

	if err := db.Create(&user).Error; err != nil {
		return c.Status(500).SendString("Failed to create user")
	}
	return c.Status(201).JSON(user)
}

3. Generate Swagger Documentation
Add comments above your handler functions and models to annotate them with Swagger-specific comments. For example:

@Summary describes what the function does.
@Description provides further explanation.
@Success defines the success response format.
@Router defines the URL route and the method (GET, POST, etc.).
Once the annotations are in place, you can generate the Swagger documentation by running:

swag init
PS C:\jsj\Exercise\8-20250112-FiberGinMVCQuasar\backend> C:\Users\josejao.INSULARNET\go\bin\swag init

http://localhost:3000/swagger/index.html
