package validate

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/lest-go/validate/utils"
)

// validateCNPJ validates a Brazilian CNPJ.
func validateCNPJ(fl validator.FieldLevel) bool {
	cnpj := fl.Field().String()
	cnpj = utils.RemoveSpecialCharacters(cnpj)

	if len(cnpj) != 14 {
		return false
	}

	if allEqual(cnpj) {
		return false
	}

	cnpjWithoutDigital := utils.RemoveDigital(cnpj)
	cnpjWithFirstDigital := getCalculateFirstDigitalCnpj(cnpjWithoutDigital)
	cnpjWithSecondDigital := getCalculateSecondDigitalCnpj(cnpjWithFirstDigital)

	return cnpjWithSecondDigital == cnpj
}

// AllEqual checks if all characters in the string are equal.
func allEqual(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] != s[0] {
			return false
		}
	}
	return true
}

func getCalculateFirstDigitalCnpj(documentWithoutDigital string) string {
	var cnpjFirstDigitTable = []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	digital := utils.GetCalculateFirstDigitalCpfCnpj(documentWithoutDigital, cnpjFirstDigitTable)
	return attach(documentWithoutDigital, digital)
}

func getCalculateSecondDigitalCnpj(documentWithoutDigital string) string {
	var cnpjSecondDigitTable = []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	digital := utils.GetCalculateFirstDigitalCpfCnpj(documentWithoutDigital, cnpjSecondDigitTable)
	return attach(documentWithoutDigital, digital)
}

func attach(cnpjWithoutDigital string, digital int) string {
	return fmt.Sprintf("%s%d", cnpjWithoutDigital, digital)
}
