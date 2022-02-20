package middlewares

import (
	"github.com/andynur/fiber-boilerplate/app/auth"
	"github.com/gofiber/fiber/v2"
)

func LimitPhoneNumbersPerRequest(c *fiber.Ctx) error {
	if auth.IsLoggedIn(c) {
		return c.Redirect("/")
	}
	return c.Next()
}
