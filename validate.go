package validate

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/lest-go/validate/utils"
	"time"
)

var validate *validator.Validate

type errorResponse struct {
	Timestamp time.Time         `json:"timestamp"`
	Error     string            `json:"error"`
	Fields    []fieldValidation `json:"validations"`
}

type fieldValidation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *errorResponse) AppendField(fe validator.FieldError) {
	message := utils.Error(fe)
	e.Fields = append(e.Fields, fieldValidation{
		Field:   fe.Field(),
		Message: message,
	})
}

func toJson(validationErrors validator.ValidationErrors) errorResponse {
	errResp := errorResponse{
		Timestamp: time.Now(),
		Error:     "StatusBadRequest",
	}

	for _, fieldError := range validationErrors {
		errResp.AppendField(fieldError)
	}
	return errResp
}

func validates(strc interface{}) (json errorResponse, err error) {
	validate = validator.New(validator.WithRequiredStructEnabled())
	err = validate.RegisterValidation("cpf", cpfValidator)
	if err != nil {
		return json, err
	}

	err = validate.Struct(strc)

	if err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if ok {
			for _, validationError := range validationErrors {
				fmt.Println(validationError.Field(), validationError.Tag(), validationError.Param())
			}
			return toJson(validationErrors), nil
		} else {
			return json, err
		}
	}
	return json, nil
}
