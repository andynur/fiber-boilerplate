package middlewares

import (
	"github.com/andynur/fiber-boilerplate/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
)

func ValidateApiRegisterPost(c *fiber.Ctx) error {
	var register models.RegisterForm
	if err := c.BodyParser(&register); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	v := validate.Struct(register)
	if !v.Validate() {
		return c.Status(401).JSON(fiber.Map{
			"error":   true,
			"message": v.Errors,
		})
	}
	c.Locals("register", register)
	return c.Next()
}
