package validate

import (
	"github.com/go-playground/validator/v10"
	"github.com/lest-go/validate/utils"
	"time"
)

var validate *validator.Validate

type ValidateI interface {
	ValidateAll(strt interface{}) (messageValidate Validation, err error)
}

type Service struct {
}

func NewValidate() ValidateI {
	return &Service{}
}

type Validation struct {
	Timestamp time.Time          `json:"timestamp"`
	Error     string             `json:"error"`
	Fields    []*FieldValidation `json:"validations"`
}

type FieldValidation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (s *Service) ValidateAll(strt interface{}) (messageValidate Validation, err error) {
	validate = validator.New(validator.WithRequiredStructEnabled())

	err = validate.RegisterValidation("cpf", cpfValidator)
	err = validate.RegisterValidation("cnpj", validateCNPJ)
	if err != nil {
		return buildMessageValidate(nil), err
	}

	err = validate.Struct(strt)

	if err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if ok {
			return buildMessageValidate(validationErrors), nil
		} else {
			return buildMessageValidate(nil), err
		}
	}

	return buildMessageValidate(nil), nil
}

func buildMessageValidate(validationErrors validator.ValidationErrors) Validation {
	if validationErrors != nil {
		errResp := Validation{
			Timestamp: time.Now(),
			Error:     "StatusBadRequest",
			Fields:    make([]*FieldValidation, 0),
		}

		for _, fieldError := range validationErrors {
			errResp.AppendField(fieldError)
		}
		return errResp
	}

	return Validation{
		Fields: make([]*FieldValidation, 0),
	}
}

func (v *Validation) AppendField(fe validator.FieldError) {
	message := utils.Error(fe)

	v.Fields = append(v.Fields, &FieldValidation{
		Field:   fe.Field(),
		Message: message,
	})
}
