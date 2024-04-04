package validate

import (
	"reflect"
	"testing"
)

func TestFieldValidationBuilder(t *testing.T) {

	builder := NewFieldValidationBuilder()
	fieldValidation := builder.
		Field("Name").
		Message("Name is required").
		Build()

	expected := &FieldValidation{
		Field:   "Name",
		Message: "Name is required",
	}

	if !reflect.DeepEqual(fieldValidation, expected) {
		t.Errorf("Expected %+v, got %+v", expected, fieldValidation)
	}

}
