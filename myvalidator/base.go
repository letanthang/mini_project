package myvalidator

import (
	validator "gopkg.in/go-playground/validator.v9"
)

// NewValidator ..
func NewValidator() *MyValidator {
	return &MyValidator{validator: validator.New()}
}

// MyValidator ..
type MyValidator struct {
	validator *validator.Validate
}

// Validate ..
func (cv *MyValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
