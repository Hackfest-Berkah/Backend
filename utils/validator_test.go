package utils

import "testing"

func TestEmailValidator(t *testing.T) {
	// Test case 1: Valid email
	validEmail := "user@example.com"
	if !EmailValidator(validEmail) {
		t.Errorf("Expected %s to be a valid email, but got invalid", validEmail)
	}

	// Test case 2: Invalid email (missing "@")
	invalidEmail := "invalidemail.com"
	if EmailValidator(invalidEmail) {
		t.Errorf("Expected %s to be an invalid email, but got valid", invalidEmail)
	}

	// Add more test cases as needed
}

func TestPasswordValidator(t *testing.T) {
	// Test case 1: Valid password
	validPassword := "SecurePass123"
	err := PasswordValidator(validPassword)
	if err != nil {
		t.Errorf("Expected %s to be a valid password, but got validation error: %s", validPassword, err)
	}

	// Test case 2: Invalid password (length less than 8)
	invalidPasswordShort := "Short1"
	err = PasswordValidator(invalidPasswordShort)
	if err == nil {
		t.Errorf("Expected %s to be an invalid password, but got no validation error", invalidPasswordShort)
	}

	// Test case 3: Invalid password (no lowercase letter)
	invalidPasswordNoLower := "UPPERCASE123"
	err = PasswordValidator(invalidPasswordNoLower)
	if err == nil {
		t.Errorf("Expected %s to be an invalid password, but got no validation error", invalidPasswordNoLower)
	}

	// Test case 4: Invalid password (no uppercase letter)
	invalidPasswordNoUpper := "lowercase123"
	err = PasswordValidator(invalidPasswordNoUpper)
	if err == nil {
		t.Errorf("Expected %s to be an invalid password, but got no validation error", invalidPasswordNoUpper)
	}

	// Test case 5: Invalid password (no digit)
	invalidPasswordNoDigit := "NoDigitHere"
	err = PasswordValidator(invalidPasswordNoDigit)
	if err == nil {
		t.Errorf("Expected %s to be an invalid password, but got no validation error", invalidPasswordNoDigit)
	}

	// Add more test cases as needed
}
