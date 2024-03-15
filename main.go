package main

import (
	"be-pemilu/migration"
	"be-pemilu/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	f := fiber.New()

	// database.ConnectDB()
	migration.Migrate()

	routes.RouteApp(f)

	// start server
	f.Listen(":8000")

}
