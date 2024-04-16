package validate

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type Validate interface {
	Struct(strt interface{}) (messageValidate []FieldValidation, err error)
}

type Service struct {
}

func NewValidate() Validate {
	return &Service{}
}

func (s *Service) Struct(strt interface{}) (messageValidate []FieldValidation, err error) {
	validate = validator.New(validator.WithRequiredStructEnabled())

	err = validate.RegisterValidation("cpf", validateCPF)
	err = validate.RegisterValidation("cnpj", validateCNPJ)
	err = validate.RegisterValidation("blank", noBlankValidator)
	if err != nil {
		return buildMessageValidate(nil), err
	}

	validation := validate.Struct(strt)

	if validation != nil {
		validationErrors, ok := validation.(validator.ValidationErrors)
		if ok {
			return buildMessageValidate(validationErrors), err
		} else {
			return buildMessageValidate(validationErrors), err
		}
	}

	return buildMessageValidate(nil), nil
}

func buildMessageValidate(validationErrors validator.ValidationErrors) []FieldValidation {
	validationFields := make([]FieldValidation, 0)
	builder := NewFieldValidationBuilder()
	for _, fieldError := range validationErrors {
		validationFields = append(validationFields, *builder.
			Field(fieldError.Field()).
			Message(MessageValidation(fieldError)).
			Build())
	}
	return validationFields
}
