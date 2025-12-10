package ssh_dto

import "github.com/go-playground/validator/v10"

type ImplementDTO struct {
	Branch ENUM_BRANCH `json:"branch" validate:"branch_name"`
}

type ENUM_BRANCH string

const (
	main ENUM_BRANCH = "main"
	test ENUM_BRANCH = "testing-new-mutex"
)

// Validação customizada
func ValidateEnumBranch(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	switch ENUM_BRANCH(value) {
	case main, test:
		return true
	default:
		return false
	}
}
