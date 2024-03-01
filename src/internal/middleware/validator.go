package middleware

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/rulanugrh/eirene/src/helper"
)

type IValidate interface {
	Validate(data interface{}) error
	ValidationMessage(err error) error
}

type Validate struct {
	validate *validator.Validate
}

func NewValidation() IValidate {
	return &Validate{
		validate: validator.New(),
	}
}

func (v *Validate) Validate(data interface{}) error {
	err := v.validate.Struct(data)
	if err != nil {
		return err
	}

	return nil
}

func (v *Validate) ValidationMessage(err error) error {
	var msg string
	for _, e := range err.(validator.ValidationErrors) {
		switch e.Tag() {
		case "required":
			msg = fmt.Sprintf("%s is required", e.Field())
		case "min":
			msg = fmt.Sprintf("%s is to short", e.Field())
		case "email":
			msg = fmt.Sprintf("%s format must email", e.Field())
		}
	}

	return helper.BadRequest(msg)
}
