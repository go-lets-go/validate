package validate

import (
	"github.com/go-playground/validator/v10"
	"github.com/lest-go/validate/utils"
)

// cpfValidator validates a Brazilian CPF.
func cpfValidator(fl validator.FieldLevel) bool {
	cpf := fl.Field().String()
	cpf = utils.RemoveSpecialCharacters(cpf)

	if len(cpf) != 11 {
		return false
	}

	if allEqual(cpf) {
		return false
	}

	cpfWithoutDigital := utils.RemoveDigital(cpf)
	cpfWithFirstDigital := getCalculateFirstDigitalCpf(cpfWithoutDigital)
	cpfWithSecondDigital := getCalculateSecondDigitalCpf(cpfWithFirstDigital)

	return cpfWithSecondDigital == cpf
}

func getCalculateFirstDigitalCpf(documentWithoutDigital string) string {
	var cpfFirstDigitTable = []int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	digital := utils.GetCalculateFirstDigitalCpfCnpj(documentWithoutDigital, cpfFirstDigitTable)
	return attach(documentWithoutDigital, digital)
}

func getCalculateSecondDigitalCpf(documentWithoutDigital string) string {
	var cpfSecondDigitTable = []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
	digital := utils.GetCalculateFirstDigitalCpfCnpj(documentWithoutDigital, cpfSecondDigitTable)
	return attach(documentWithoutDigital, digital)
}
