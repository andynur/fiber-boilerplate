package routes

import (
	"github.com/andynur/fiber-boilerplate/app"
	"github.com/andynur/fiber-boilerplate/app/controllers"
	"github.com/andynur/fiber-boilerplate/app/middlewares"
	"github.com/gofiber/fiber/v2"
)

func LandingRoutes(web fiber.Router) {
	web.Get("/", controllers.Landing)
	web.Get("/ping", Pong)
	web.Get("/all-routes", AllRoutes)
	web.Get("/do/verify-email", middlewares.ValidateConfirmToken, controllers.VerifyRegisteredEmail)
}

func Pong(c *fiber.Ctx) error {
	return c.SendString("Pong")
}

func AllRoutes(c *fiber.Ctx) error {
	return c.JSON(app.Http.Server.Stack())
}
