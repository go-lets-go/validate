package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

const (
	FIELD_MSG_DEFAULT = "Validation: Field '%s' failed on the '%s = %s'"

	FIELD_MSG_MIN = "Validation: Field '%s' failed on the %s = '%s'"
	FIELD_MSG_MAX = "Validation: Field '%s' failed on the '%s = %s'"

	FIELD_MSG_GREATE_THEN = "Validation: Field '%s' failed greater than or equal = '%s'"
	FIELD_MSG_LESS_THEN   = "Validation: Field '%s' failed less than or equal = '%s'"
	FIELD_MSG_CHOOSE      = "Validation: Field '%s' failed when choosing between '%s'"
	FIELD_MSG_EMAIL       = "Validation: Field '%s' is invalid"

	PRE_MSG = "Validation:"
)

func Error(fe validator.FieldError) string {
	switch fe.Tag() {
	case "email":
		return validatorEmail(fe)
	case "gte":
		return validatorGreateEquals(fe)
	case "lte":
		return validatorLessEquals(fe)
	case "min":
		return validatorMin(fe)
	case "max":
		return validatorMax(fe)
	case "cpf":
		return validatorCpf(fe)
	default:
		return validatorDefault(fe)
	}

}

func validatorGreateEquals(fe validator.FieldError) string {
	return fmt.Sprintf(FIELD_MSG_GREATE_THEN, fe.Field(), fe.Param())
}

func validatorLessEquals(fe validator.FieldError) string {
	return fmt.Sprintf(FIELD_MSG_LESS_THEN, fe.Field(), fe.Param())
}

func validatorMax(fe validator.FieldError) string {
	return fmt.Sprintf(FIELD_MSG_MIN, fe.Field(), fe.Tag(), fe.Param())
}

func validatorMin(fe validator.FieldError) string {
	return fmt.Sprintf(FIELD_MSG_MIN, fe.Field(), fe.Tag(), fe.Param())
}

func validatorCpf(fe validator.FieldError) string {
	return fmt.Sprintf("%s Field '%s' is invalid", PRE_MSG, fe.Field())
}

func validatorEmail(fe validator.FieldError) string {
	return fmt.Sprintf(FIELD_MSG_EMAIL, fe.Field())
}

func validatorDefault(fe validator.FieldError) string {
	return fmt.Sprintf(FIELD_MSG_DEFAULT, fe.Field(), fe.Tag(), fe.Param())
}
