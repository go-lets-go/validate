package utils

import (
	"regexp"
	"strconv"
)

func RemoveDigital(cpfCnpj string) string {
	if len(cpfCnpj) == 11 {
		return cpfCnpj[0:9]
	}
	return cpfCnpj[:12]
}

func GetCalculateFirstDigitalCpfCnpj(documentWithoutDigital string, DigitTable []int) int {
	sum := sumDigit(documentWithoutDigital, DigitTable)
	rest := sum % 11

	digital := 0
	if rest >= 2 {
		digital = 11 - rest
	}
	return digital
}

func sumDigit(s string, table []int) int {
	if len(s) != len(table) {
		return 0
	}

	sum := 0
	for i, v := range table {
		c := string(s[i])
		d, err := strconv.Atoi(c)
		if err == nil {
			sum += v * d
		}
	}

	return sum
}

// RemoveSpecialCharacters remove special characters
func RemoveSpecialCharacters(s string) string {
	regex := regexp.MustCompile(`[^0-9]`)
	return regex.ReplaceAllString(s, "")
}
