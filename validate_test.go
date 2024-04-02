package validate

import (
	"testing"
)

type Person struct {
	Name   string `validate:"required,min=3,max=100"`
	Age    int    `validate:"gte=0,lte=120"`
	Email  string `validate:"required,email"`
	Gender string `validate:"oneof=male female"`
	CPF    string `validate:"required,cpf"`
	Company
}

type Company struct {
	Name string `json:"name" validate:"required"`
	CNPJ string `json:"cnpj" validate:"required,cnpj"`
}

func Test_validates(t *testing.T) {
	// Teste para um cenário bem-sucedido
	t.Run("Success", func(t *testing.T) {

		company := Company{
			Name: "techcomp",
			CNPJ: "41.756.914/0001-64",
		}

		person := Person{
			Name:    "John Doe",
			Age:     30,
			Email:   "jamersom@gmail.com",
			Gender:  "male",
			CPF:     "324.209.740-85",
			Company: company,
		}

		validate := NewValidate()

		result, err := validate.ValidateAll(person)
		if err != nil {
			t.Errorf("expected no error, got: %v", err)
		}

		if len(result.Fields) != 0 {
			t.Errorf("expected no validation errors")
		}
	})

	//// Teste para validação de CPF
	//t.Run("cpf invalid", func(t *testing.T) {
	//	person := Person{
	//		Name:   "John Doe",
	//		Age:    30,
	//		Email:  "jamersom@gmail.com",
	//		Gender: "male",
	//		CPF:    "00011122233",
	//	}
	//
	//	result, err := ValidateAll(person)
	//	if err != nil {
	//		t.Errorf("expected no error, got %v", err)
	//	}
	//
	//	expected := "Validation: Field 'CPF' is invalid"
	//
	//	for _, validation := range result.Fields {
	//		if validation.Message != expected {
	//			t.Errorf("expected '%s' got '%s'", expected, validation.Message)
	//		}
	//	}
	//})
	//
	//t.Run("cpf required", func(t *testing.T) {
	//	person := Person{
	//		Name:   "John Doe",
	//		Age:    30,
	//		Email:  "jamersom@gmail.com",
	//		Gender: "male",
	//		CPF:    "", // CPF inválido
	//	}
	//
	//	result, err := ValidateAll(person)
	//	if err != nil {
	//		t.Errorf("expected no error, got %v", err)
	//	}
	//
	//	expected := "Validation: Field 'CPF' is required"
	//
	//	for _, validation := range result.Fields {
	//		if validation.Message != expected {
	//			t.Errorf("expected '%s' got '%s'", expected, validation.Message)
	//		}
	//	}
	//
	//})
}
