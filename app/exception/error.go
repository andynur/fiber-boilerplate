package exception

import (
	"github.com/andynur/fiber-boilerplate/app/data/serializer"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ErrorPanic(err interface{}) {
	if err != nil {
		panic(err)
	}
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	_, ok := err.(ValidationError)
	if ok {
		// var obj interface{}
		// _ = json.Unmarshal([]byte(err.Error()), &obj)
		return ctx.Status(400).JSON(serializer.ResponseError{
			Status:  400,
			Code:    "E_VALIDATION_ERROR",
			Data:    struct{}{},
			Message: err.Error(),
		})
	}

	if err == gorm.ErrRecordNotFound {
		return ctx.Status(404).JSON(serializer.ResponseError{
			Status:  404,
			Code:    "E_RECORD_NOT_FOUND",
			Data:    nil,
			Message: err.Error(),
		})
	}

	if err.Error() == "E_PHONE_REGISTERED" {
		return ctx.Status(400).JSON(serializer.ResponseError{
			Status:  400,
			Code:    "E_PHONE_REGISTERED",
			Data:    nil,
			Message: "phone number already registered",
		})
	}

	if err.Error() == "E_USERNAME_REGISTERED" {
		return ctx.Status(400).JSON(serializer.ResponseError{
			Status:  400,
			Code:    "E_USERNAME_REGISTERED",
			Data:    nil,
			Message: "username already registered",
		})
	}

	if err.Error() == "E_EMAIL_REGISTERED" {
		return ctx.Status(400).JSON(serializer.ResponseError{
			Status:  400,
			Code:    "E_EMAIL_REGISTERED",
			Data:    nil,
			Message: "email already registered",
		})
	}

	if err.Error() == "E_PASSWORD_OLD_NOTMATCH" {
		return ctx.Status(400).JSON(serializer.ResponseError{
			Status:  400,
			Code:    "E_PASSWORD_OLD_NOTMATCH",
			Data:    nil,
			Message: "old password not match",
		})
	}

	return ctx.Status(500).JSON(serializer.ResponseError{
		Status: 500,
		Code:   "E_INTERNAL_SERVER_ERROR",
		Data:   err.Error(),
	})
}
