package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

const (
	FIELD_MSG_REQUIRED     = "Validation: Field '%s' is required"
	FIELD_MSG_EMAIL        = "Validation: Field '%s' is invalid"
	FIELD_MSG_CPF_AND_CNPJ = "Validation: Field '%s' is invalid"

	FIELD_MSG_MIN = "Validation: Field '%s' failed on the %s = '%s'"
	FIELD_MSG_MAX = "Validation: Field '%s' failed on the '%s = %s'"

	FIELD_MSG_GREATE_THEN = "Validation: Field '%s' failed greater than or equal = '%s'"
	FIELD_MSG_LESS_THEN   = "Validation: Field '%s' failed less than or equal = '%s'"

	//FIELD_MSG_CHOOSE      = "Validation: Field '%s' failed when choosing between '%s'"
	FIELD_MSG_DEFAULT = "Validation: Field '%s' failed on the '%s = %s'"
)

func Error(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return messageRequired(fe)
	case "email":
		return messageEmail(fe)
	case "gte":
		return messageGreateEquals(fe)
	case "lte":
		return messageLessEquals(fe)
	case "min":
		return messageMin(fe)
	case "max":
		return messageMax(fe)
	case "cpf":
		return messageCpfOrCnpj(fe)
	case "cnpj":
		return messageCpfOrCnpj(fe)
	default:
		return messageDefault(fe)
	}

}

func messageRequired(fe validator.FieldError) string {
	return fmt.Sprintf(FIELD_MSG_REQUIRED, fe.Field())
}

func messageEmail(fe validator.FieldError) string {
	return fmt.Sprintf(FIELD_MSG_EMAIL, fe.Field())
}

func messageGreateEquals(fe validator.FieldError) string {
	return fmt.Sprintf(FIELD_MSG_GREATE_THEN, fe.Field(), fe.Param())
}

func messageLessEquals(fe validator.FieldError) string {
	return fmt.Sprintf(FIELD_MSG_LESS_THEN, fe.Field(), fe.Param())
}

func messageMax(fe validator.FieldError) string {
	return fmt.Sprintf(FIELD_MSG_MAX, fe.Field(), fe.Tag(), fe.Param())
}

func messageMin(fe validator.FieldError) string {
	return fmt.Sprintf(FIELD_MSG_MIN, fe.Field(), fe.Tag(), fe.Param())
}

func messageCpfOrCnpj(fe validator.FieldError) string {
	return fmt.Sprintf(FIELD_MSG_CPF_AND_CNPJ, fe.Field())
}

func messageDefault(fe validator.FieldError) string {
	return fmt.Sprintf(FIELD_MSG_DEFAULT, fe.Field(), fe.Tag(), fe.Param())
}
