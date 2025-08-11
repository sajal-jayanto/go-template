package utils

import (
	"errors"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func FormatValidationErrors(err error) []string {
	var errors []string
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			errors = append(errors, ValidationErrorToText(e))
		}
	} else if err != nil {
		errors = append(errors, err.Error())
	}
	return errors
}

func GetParamInt(ctx *fiber.Ctx, param string) (int, error) {
	idStr := ctx.Params(param)
	if idStr == "" {
		return 0, errors.New("missing " + param + " parameter")
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, errors.New("invalid " + param + " parameter")
	}
	return id, nil
}

func ValidationErrorToText(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " is required"
	case "min":
		return fe.Field() + " must be at least " + fe.Param() + " characters long"
	case "max":
		return fe.Field() + " must be at most " + fe.Param() + " characters long"
	case "email":
		return fe.Field() + " must be a valid email address"
	default:
		return fe.Field() + " is invalid"
	}
}