package validate

import (
	"testing"
)

type Person struct {
	Name  string `validate:"required,blank,min=3,max=10"`
	Email string `validate:"email"`
	CPF   string `validate:"cpf"`
	Company
}

type Company struct {
	CNPJ string `validate:"cnpj"`
}

func Test_validate(t *testing.T) {
	// INITIATE - FIELD NAME
	t.Run("validate blank tag valid ( )", func(t *testing.T) {
		results, err := NewValidate().ValidateAll(Person{Name: " "})
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

	t.Run("validate blank tag invalid '      '", func(t *testing.T) {
		results, err := NewValidate().ValidateAll(Person{Name: "     "})
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

	t.Run("validate blank tag valid (Jamerson)", func(t *testing.T) {
		results, err := NewValidate().ValidateAll(Person{Name: "Jamerson"})
		expected := "Validation: Field 'Name' failed is blank"
		if err != nil {
			t.Error("expected error, got nil")
		} else {
			for _, result := range results {
				if result.Field != "Name" {
					continue
				} else {
					if result.Message == expected {
						t.Errorf("expected error: %s, got: %v", expected, results)
					}
				}
			}
		}
	})
	// END - FIELD NAME

	// INITIATE - CPF
	t.Run("validate cpf tag invalid ()", func(t *testing.T) {
		results, err := NewValidate().ValidateAll(Person{CPF: ""})
		expected := "Validation: Field 'CPF' is invalid"
		if err != nil {
			t.Error("expected error, got nil")
		} else {
			for _, result := range results {
				if result.Field == "CPF" && result.Message == expected {
					return
				}
			}
			t.Errorf("expected error: %s, got: %v", expected, results)
		}
	})

	t.Run("validate cpf tag invalid (   .   .   -  )", func(t *testing.T) {
		results, err := NewValidate().ValidateAll(Person{CPF: "   .   .   -  "})
		expected := "Validation: Field 'CPF' is invalid"
		if err != nil {
			t.Error("expected error, got nil")
		} else {
			for _, result := range results {
				if result.Field == "CPF" && result.Message == expected {
					return
				}
			}
			t.Errorf("expected error: %s, got: %v", expected, results)
		}
	})

	t.Run("validate cpf tag invalid (324.209.740-)", func(t *testing.T) {
		results, err := NewValidate().ValidateAll(Person{CPF: "324.209.740-"})
		expected := "Validation: Field 'CPF' is invalid"
		if err != nil {
			t.Error("expected error, got nil")
		} else {
			for _, result := range results {
				if result.Field == "CPF" && result.Message == expected {
					return
				}
			}
			t.Errorf("expected error: %s, got: %v", expected, results)
		}
	})

	t.Run("validate cpf tag invalid (324.209.740-96)", func(t *testing.T) {
		results, err := NewValidate().ValidateAll(Person{CPF: "324.209.740-96"})
		expected := "Validation: Field 'CPF' is invalid"
		if err != nil {
			t.Error("expected error, got nil")
		} else {
			for _, result := range results {
				if result.Field == "CPF" && result.Message == expected {
					return
				}
			}
			t.Errorf("expected error: %s, got: %v", expected, results)
		}
	})

	t.Run("validate cpf tag invalid (aaa.bbb.ccc-dd)", func(t *testing.T) {
		results, err := NewValidate().ValidateAll(Person{CPF: "aaa.bbb.ccc-dd"})
		expected := "Validation: Field 'CPF' is invalid"
		if err != nil {
			t.Error("expected error, got nil")
		} else {
			for _, result := range results {
				if result.Field == "CPF" && result.Message == expected {
					return
				}
			}
			t.Errorf("expected error: %s, got: %v", expected, results)
		}
	})

	t.Run("validate cpf tag success (324.209.740-85)", func(t *testing.T) {
		results, err := NewValidate().ValidateAll(Person{CPF: "324.209.740-85"})
		expected := "Validation: Field 'CPF' is invalid"
		if err != nil {
			t.Error("expected error, got nil")
		} else {
			for _, result := range results {
				if result.Field != "CPF" {
					continue
				} else {
					if result.Message == expected {
						t.Errorf("expected error: %s, got: %v", expected, results)
					}
				}
			}
		}
	})
	// END - CPF

	// INITIATE - CNPJ
	t.Run("validate cnpj tag invalid ()", func(t *testing.T) {
		results, err := NewValidate().ValidateAll(Company{CNPJ: ""})
		expected := "Validation: Field 'CNPJ' is invalid"
		if err != nil {
			t.Error("expected error, got nil")
		} else {
			for _, result := range results {
				if result.Field == "CNPJ" && result.Message == expected {
					return
				}
			}
			t.Errorf("expected error: %s, got: %v", expected, results)
		}
	})

	t.Run("validate cnpj tag invalid (  .   .   /    -  )", func(t *testing.T) {
		results, err := NewValidate().ValidateAll(Company{CNPJ: "  .   .   /    -  "})
		expected := "Validation: Field 'CNPJ' is invalid"
		if err != nil {
			t.Error("expected error, got nil")
		} else {
			for _, result := range results {
				if result.Field == "CNPJ" && result.Message == expected {
					return
				}
			}
			t.Errorf("expected error: %s, got: %v", expected, results)
		}
	})

	t.Run("validate cnpj tag invalid (41.756.914/0001-)", func(t *testing.T) {
		results, err := NewValidate().ValidateAll(Company{CNPJ: "41.756.914/0001-"})
		expected := "Validation: Field 'CNPJ' is invalid"
		if err != nil {
			t.Error("expected error, got nil")
		} else {
			for _, result := range results {
				if result.Field == "CNPJ" && result.Message == expected {
					return
				}
			}
			t.Errorf("expected error: %s, got: %v", expected, results)
		}
	})

	t.Run("validate cnpj tag invalid (41.746.914/0001-01)", func(t *testing.T) {
		results, err := NewValidate().ValidateAll(Company{CNPJ: "41.746.914/0001-01"})
		expected := "Validation: Field 'CNPJ' is invalid"
		if err != nil {
			t.Error("expected error, got nil")
		} else {
			for _, result := range results {
				if result.Field == "CNPJ" && result.Message == expected {
					return
				}
			}
			t.Errorf("expected error: %s, got: %v", expected, results)
		}
	})

	t.Run("validate cnpj tag invalid (aa.bbb.ccc/dddd-ee)", func(t *testing.T) {
		results, err := NewValidate().ValidateAll(Company{CNPJ: "aa.bbb.ccc/dddd-ee"})
		expected := "Validation: Field 'CNPJ' is invalid"
		if err != nil {
			t.Error("expected error, got nil")
		} else {
			for _, result := range results {
				if result.Field == "CNPJ" && result.Message == expected {
					return
				}
			}
			t.Errorf("expected error: %s, got: %v", expected, results)
		}
	})

	t.Run("validate cnpj tag success (41.756.914/0001-64)", func(t *testing.T) {
		results, err := NewValidate().ValidateAll(Company{CNPJ: "41.756.914/0001-64"})
		expected := "Validation: Field 'CNPJ' is invalid"
		if err != nil {
			t.Error("expected error, got nil")
		} else {
			if len(results) != 0 {
				t.Errorf("expected error: %s, got: %v", expected, results)
			}
		}
	})
	// END - CNPJ
}
