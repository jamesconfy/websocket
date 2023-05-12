package service

import "github.com/go-playground/validator/v10"

type ValidationService interface {
	Validate(any) error
}

type validationStruct struct{}

func NewValidationService() ValidationService {
	return &validationStruct{}
}

func (v *validationStruct) Validate(a any) error {
	validate := validator.New()
	return validate.Struct(a)
}
