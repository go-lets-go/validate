package validate

type FieldValidation struct {
	Field   string
	Message string
}

type ValidationBuilder struct {
	field   string
	message string
}

// NewFieldValidationBuilder creates a new ValidationBuilder
func NewFieldValidationBuilder() *ValidationBuilder {
	return &ValidationBuilder{}
}

// Field sets the field of the ValidationBuilder
func (b *ValidationBuilder) Field(field string) *ValidationBuilder {
	b.field = field
	return b
}

// Message sets the message of the ValidationBuilder
func (b *ValidationBuilder) Message(message string) *ValidationBuilder {
	b.message = message
	return b
}

// Build creates the final FieldValidation object using the values defined in the ValidationBuilder
func (b *ValidationBuilder) Build() *FieldValidation {
	return &FieldValidation{
		Field:   b.field,
		Message: b.message,
	}
}
