package services

import (
	"fmt"
	"time"

	"github.com/andynur/fiber-boilerplate/app"
	"github.com/andynur/fiber-boilerplate/utils"
	"github.com/gofiber/fiber/v2"
)

func SendPasswordResetEmail(email string, baseURL string) {
	resetEmail := fmt.Sprintf("%s-reset-%d", email, time.Now().Unix())
	resetLink := GeneratePasswordResetURL(resetEmail, baseURL)
	htmlBody := app.Http.Mail.PrepareHtml("emails/password-reset", fiber.Map{
		"reset_link": resetLink,
	})
	app.Http.Mail.Send(email, "You asked to reset? Please click here!", htmlBody, "", "")
}

func SendConfirmationEmail(email string, baseURL string) {
	confirmLink := GenerateConfirmURL(email, baseURL)
	htmlBody := app.Http.Mail.PrepareHtml("emails/confirm", fiber.Map{
		"confirm_link": confirmLink,
	})
	app.Http.Mail.Send(email, "Is it you? Please confirm!", htmlBody, "", "")
}

func GenerateConfirmURL(email string, baseURL string) string {
	token := utils.EncryptToken(email, app.Http.Server.Key)
	uri := fmt.Sprintf("%s/do/verify-email?t=%s", baseURL, token)
	return uri
}

func GeneratePasswordResetURL(email string, baseURL string) string {
	token := utils.EncryptToken(email, app.Http.Server.Key)
	uri := fmt.Sprintf("%s/reset-password?t=%s", baseURL, token)
	return uri
}
