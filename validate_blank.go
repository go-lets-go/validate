package validate

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func noBlankValidator(fl validator.FieldLevel) (isValid bool) {
	re := regexp.MustCompile(`^\s*$`)
	isValid = false
	if !re.MatchString(fl.Field().String()) {
		isValid = true
	}
	return isValid
}
