package validate

import (
	"log/slog"
	"testing"
)

type Person struct {
	Name   string `json:"name" validate:"required,min=3,max=100"`
	Age    int    `json:"age"  validate:"gte=0,lte=120"`
	Email  string `json:"email" validate:"required,email"`
	Gender string `json:"gender"  validate:"oneof=male female"`
	CPF    string `json:"cpf" validate:"cpf"`
}

func Test_validates(t *testing.T) {
	person := Person{
		Name:   "John Doe",
		Age:    30,
		Email:  "jamersom@gmail.com",
		Gender: "male",
		CPF:    "000.111.245-88",
	}

	result, err := validates(person)
	if err != nil {
		slog.Error("error validator Person:", err)
	}
	expected := "Validation: Field 'CPF' is invalid"

	for _, validation := range result.Fields {
		println(validation.Message)

		if validation.Message != expected {
			t.Errorf("expected '%s' got '%s'", expected, validation.Message)
		}
	}
}
