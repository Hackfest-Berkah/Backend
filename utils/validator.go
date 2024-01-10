package utils

import (
	"fmt"
	"regexp"
	"unicode"
)

func EmailValidator(email string) bool {
	return regexp.MustCompile(`@`).MatchString(email)
}

func PasswordValidator(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("Password must be at least 8 characters long.")
	}
	if !containsLowerCase(password) {
		return fmt.Errorf("Password must have at least one lowercase letter.")
	}
	if !containsUpperCase(password) {
		return fmt.Errorf("Password must have at least one uppercase letter.")
	}
	if !containsDigit(password) {
		return fmt.Errorf("Password must have at least one digit.")
	}
	return nil
}

func containsLowerCase(s string) bool {
	for _, char := range s {
		if unicode.IsLower(char) {
			return true
		}
	}
	return false
}

func containsUpperCase(s string) bool {
	for _, char := range s {
		if unicode.IsUpper(char) {
			return true
		}
	}
	return false
}

func containsDigit(s string) bool {
	for _, char := range s {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}
