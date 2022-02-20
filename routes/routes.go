package routes

import (
	"github.com/andynur/fiber-boilerplate/app/middlewares"
	"github.com/gofiber/fiber/v2"
)

func LoadRoutes(app *fiber.App) {
	api := app.Group("api").Use(middlewares.AuthApi())
	web := app.Group("")
	ApiRoutes(api)
	WebRoutes(web)
}
