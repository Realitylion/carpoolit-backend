package main

import (
	"carpool-backend/database"
	"carpool-backend/routes/bookings"

	// "carpool-backend/routes/rides"
	// "carpool-backend/routes/users"
	"log"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/booking", bookings.CreateBooking)
	// app.Post("/ride", rides.CreateRide)
	// app.Post("/user", users.CreateUser)

}

func main() {
	app := fiber.New()
	database.ConnectToDB()

	SetupRoutes(app)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	log.Fatal(app.Listen(":3000"))
}
