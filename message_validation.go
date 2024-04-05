package validate

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

const (
	MESSAGE_REQUIRED = "Validation: Field '%s' is required"
	MESSAGE_BLANK    = "Validation: Field '%s' failed is %s"
	MESSAGE_EMAIL    = "Validation: Field '%s' is invalid"

	MESSAGE_CPF_AND_CNPJ = "Validation: Field '%s' is invalid"

	MESSAGE_MIM_MAX = "Validation: Field '%s' failed the %s length of %s characters"

	//MESSAGE_CHOOSE      = "Validation: Field '%s' failed when choosing between '%s'"

	MESSAGE_GREATE_THEN = "Validation: Field '%s' failed greater than or equal = '%s'"
	MESSAGE_LESS_THEN   = "Validation: Field '%s' failed less than or equal = '%s'"

	MESSAGE_DEFAULT          = "Validation: Field '%s' failed on the '%s = %s'"
	MESSAGE_DEFAULT_NO_PARAM = "Validation: Field '%s' failed on the '%s'"
)

func MessageValidation(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return messageRequired(fe)
	case "blank":
		return messageBlank(fe)
	case "email":
		return messageEmail(fe)
	case "gte":
		return messageGreateEquals(fe)
	case "lte":
		return messageLessEquals(fe)
	case "min":
		return messageMinMax(fe)
	case "max":
		return messageMinMax(fe)
	case "cpf":
		return messageCpfOrCnpj(fe)
	case "cnpj":
		return messageCpfOrCnpj(fe)
	default:
		return messageDefault(fe)
	}

}

func messageRequired(fe validator.FieldError) string {
	return fmt.Sprintf(MESSAGE_REQUIRED, fe.Field())
}

func messageEmail(fe validator.FieldError) string {
	return fmt.Sprintf(MESSAGE_EMAIL, fe.Field())
}

func messageGreateEquals(fe validator.FieldError) string {
	return fmt.Sprintf(MESSAGE_GREATE_THEN, fe.Field(), fe.Param())
}

func messageLessEquals(fe validator.FieldError) string {
	return fmt.Sprintf(MESSAGE_LESS_THEN, fe.Field(), fe.Param())
}

func messageMinMax(fe validator.FieldError) string {
	return fmt.Sprintf(MESSAGE_MIM_MAX, fe.Field(), fe.Tag(), fe.Param())
}

func messageCpfOrCnpj(fe validator.FieldError) string {
	return fmt.Sprintf(MESSAGE_CPF_AND_CNPJ, strings.ToUpper(fe.Field()))
}

func messageDefault(fe validator.FieldError) string {
	if fe.Param() == "" {
		return fmt.Sprintf(MESSAGE_DEFAULT_NO_PARAM, fe.Field(), fe.Tag())
	}
	return fmt.Sprintf(MESSAGE_DEFAULT, fe.Field(), fe.Tag(), fe.Param())
}

func messageBlank(fe validator.FieldError) string {
	return fmt.Sprintf(MESSAGE_BLANK, fe.Field(), fe.Tag())
}
