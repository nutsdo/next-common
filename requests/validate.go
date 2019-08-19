package requests

import "gopkg.in/go-playground/validator.v9"

var Validate *validator.Validate

func init()  {
	Validate = NewValidate()
}

func NewValidate() *validator.Validate{
	return validator.New()
}