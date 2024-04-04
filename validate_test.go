package validate

import (
	"testing"
)

type Person struct {
	Name   string `validate:"required,blank,min=3,max=10"`
	Age    int    `validate:"gte=0,lte=120"`
	Email  string `validate:"email"`
	Gender string `validate:"oneof=male female"`
	CPF    string `validate:"required,cpf"`
	Company
}

type Company struct {
	Name string `validate:"required"`
	CNPJ string `validate:"required,cnpj"`
}

func Test_validate(t *testing.T) {
	// Teste para um cenário bem-sucedido
	t.Run("Success", func(t *testing.T) {
		person := Person{
			Name:   "Jamerson",
			Age:    30,
			Email:  "lestsgo@git.com",
			Gender: "male",
			CPF:    "324.209.740-85",
			Company: Company{
				Name: "techcomp",
				CNPJ: "41.756.914/0001-64",
			},
		}

		validate := NewValidate()
		results, err := validate.ValidateAll(person)
		for _, result := range results {
			if err != nil {
				t.Errorf("expected no error, got: %v", err)
			}

			if len(result.Message) != 0 {
				t.Errorf("expected no validation errors")
			}
		}
	})

	// INITIATE - FIELD NAME
	t.Run("validate required name fields", func(t *testing.T) {
		person := Person{
			Name:   "",
			Age:    30,
			Email:  "teste@gmail.com",
			Gender: "male",
			CPF:    "324.209.740-85",
			Company: Company{
				Name: "techcomp",
				CNPJ: "41.756.914/0001-64",
			},
		}

		validate := NewValidate()
		results, err := validate.ValidateAll(person)
		if err != nil {
			t.Error("expected error, got nil")
		} else {
			expected := "Validation: Field 'Name' is required"
			for _, result := range results {
				if result.Field == "Name" && result.Message == expected {
					return
				}
			}
			t.Errorf("expected error: %s, got: %v", expected, results)
		}
	})

	t.Run("validate blank name fields", func(t *testing.T) {
		person := Person{
			Name:   " ",
			Age:    30,
			Email:  "teste@gmail.com",
			Gender: "male",
			CPF:    "324.209.740-85",
			Company: Company{
				Name: "techcomp",
				CNPJ: "41.756.914/0001-64",
			},
		}

		results, err := NewValidate().ValidateAll(person)
		expected := "Validation: Field 'Name' failed is blank"
		if err != nil {
			t.Error("expected error, got nil")
		} else {
			for _, result := range results {
				if result.Field == "Name" && result.Message == expected {
					return
				}
			}
			t.Errorf("expected error: %s, got: %v", expected, results)
		}
	})

	t.Run("validate name with less than 3 characters", func(t *testing.T) {
		person := Person{
			Name:   "Ja",
			Age:    30,
			Email:  "teste@gmail.com",
			Gender: "male",
			CPF:    "324.209.740-85",
			Company: Company{
				Name: "techcomp",
				CNPJ: "41.756.914/0001-64",
			},
		}

		validate := NewValidate()
		results, err := validate.ValidateAll(person)
		if err != nil {
			t.Error("expected error, got nil")
		} else {
			expected := "Validation: Field 'Name' failed the min length of 3 characters"
			for _, result := range results {
				if result.Field == "Name" && result.Message == expected {
					return
				}
			}
			t.Errorf("expected error: %s, got: %v", expected, results)
		}
	})

	t.Run("validate name with greater than 10 characters", func(t *testing.T) {
		person := Person{
			Name:   "Jamerson Vasconcelos",
			Age:    30,
			Email:  "teste@gmail.com",
			Gender: "male",
			CPF:    "324.209.740-85",
			Company: Company{
				Name: "techcomp",
				CNPJ: "41.756.914/0001-64",
			},
		}

		validate := NewValidate()
		results, err := validate.ValidateAll(person)
		if err != nil {
			t.Error("expected error, got nil")
		} else {
			expected := "Validation: Field 'Name' failed the max length of 10 characters"
			for _, result := range results {
				if result.Field == "Name" && result.Message == expected {
					return
				}
			}
			t.Errorf("expected error: %s, got: %v", expected, results)
		}
	})
	// FIM - FIELD NAME

	t.Run("validate email with blank value", func(t *testing.T) {
		person := Person{
			Name:   "Jamerson",
			Age:    30,
			Email:  "",
			Gender: "male",
			CPF:    "324.209.740-85",
			Company: Company{
				Name: "techcomp",
				CNPJ: "41.756.914/0001-64",
			},
		}

		results, err := NewValidate().ValidateAll(person)
		expected := "Validation: Field 'Email' is invalid"
		if err != nil {
			t.Error("expected error, got nil")
		} else {
			for _, result := range results {
				if result.Field == "Email" && result.Message == expected {
					return
				}
			}
			t.Errorf("expected error: %s, got: %v", expected, results)
		}
	})

	t.Run("validate email without first part", func(t *testing.T) {
		person := Person{
			Name:   "Jamerson",
			Age:    30,
			Email:  "@goletsgo.com",
			Gender: "male",
			CPF:    "324.209.740-8",
			Company: Company{
				Name: "techcomp",
				CNPJ: "41.756.914/0001-64",
			},
		}

		results, err := NewValidate().ValidateAll(person)
		expected := "Validation: Field 'Email' is invalid"
		if err != nil {
			t.Error("expected error, got nil")
		} else {
			for _, result := range results {
				if result.Field == "Email" && result.Message == expected {
					return
				}
			}
			t.Errorf("expected error: %s, got: %v", expected, results)
		}
	})

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
}
