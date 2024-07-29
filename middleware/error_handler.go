package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CustomErrorHandler(c *fiber.Ctx, errs error) error {
	if validationErrors, ok := errs.(validator.ValidationErrors); ok {
		errMsgs := make([]string, 0)
		for _, err := range validationErrors {
			switch err.Tag() {
			case "required":
				errMsgs = append(errMsgs, fmt.Sprintf("'%v' is a required field", err.Field()))
			case "url":
				errMsgs = append(errMsgs, fmt.Sprintf("'%v' is not a valid URL", err.Value()))
			default:
				errMsgs = append(errMsgs, fmt.Sprintf("'%v' is invalid", err.Field()))
			}
		}

		return c.Status(http.StatusBadRequest).SendString(strings.Join(errMsgs, " and "))
	}

	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(errs, &e) {
		code = e.Code
	}

	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	return c.Status(code).SendString(errs.Error())
}
