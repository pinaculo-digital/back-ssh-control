package validation

import (
	"errors"
	"regexp"
	"strings"
)

var (
	errInvalidEmail         = errors.New("invalid email format")
	errInvalidCPF           = errors.New("invalid CPF format")
	errInvalidCNPJ          = errors.New("invalid CNPJ format")
	errWeakPassword         = errors.New("password must be at least 8 characters long, contain uppercase, lowercase, number, and special character")
	errInvalidPinaculoEmail = errors.New("email must have domain pinaculodigital.com.br")
)

func IsValidEmail(email string) error {
	if email == "" {
		return errInvalidEmail
	}

	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(pattern)
	if !re.MatchString(email) {
		return errInvalidEmail
	}

	return nil
}

func IsValidCPF(cpf string) error {
	// Remove non-digits
	cpf = regexp.MustCompile(`[^\d]`).ReplaceAllString(cpf, "")

	if len(cpf) != 11 {
		return errInvalidCPF
	}

	return nil
}

func IsValidCNPJ(cnpj string) error {
	cnpj = regexp.MustCompile(`[^\d]`).ReplaceAllString(cnpj, "")

	if len(cnpj) != 14 {
		return errInvalidCNPJ
	}
	return nil
}

func IsStrongPassword(password string) error {

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`).MatchString(password)

	if !hasUpper || !hasLower || !hasNumber || !hasSpecial {
		return errWeakPassword
	}

	return nil
}

func IsValidEmailPinaculo(email string) error {
	err := IsValidEmail(email)
	if err != nil {
		return err
	}

	// Check if domain is pinaculodigital.com.br
	if !strings.HasSuffix(strings.ToLower(email), "@pinaculodigital.com.br") {
		return errInvalidPinaculoEmail
	}

	return nil
}
