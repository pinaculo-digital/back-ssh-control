package auth_dto

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

// Custom validator para domínio específico
func ValidatePinaculoDomain(fl validator.FieldLevel) bool {
	email := fl.Field().String()

	// Verifica se é um email válido
	if !strings.Contains(email, "@") {
		return false
	}

	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}

	domain := parts[1]

	allowedDomains := []string{
		"pinaculodigital.com.br",
	}

	for _, allowed := range allowedDomains {
		if domain == allowed {
			return true
		}
	}
	return false
}
