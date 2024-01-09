package utils

import (
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
	"time"
)

func TestStringToInteger(t *testing.T) {
	// Create a gin.Context for testing with gin.Default()
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	// Test valid input
	result := StringToInteger("123", c)
	if result != 123 {
		t.Errorf("Expected: 123, Got: %d", result)
	}

	// Test invalid input
	result = StringToInteger("invalid", c)
	if result != 0 {
		t.Errorf("Expected: 0, Got: %d", result)
	}
	// Add more test cases as needed
}

func TestStringToFloat(t *testing.T) {
	// Create a gin.Context for testing with gin.Default()
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	// Test valid input
	result := StringToFloat("123.45", c)
	if result != 123.45 {
		t.Errorf("Expected: 123.45, Got: %f", result)
	}

	// Test invalid input
	result = StringToFloat("invalid", c)
	if result != 0.0 {
		t.Errorf("Expected: 0.0, Got: %f", result)
	}
	// Add more test cases as needed
}

func TestStringToUint(t *testing.T) {
	// Create a gin.Context for testing with gin.Default()
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	// Test valid input
	result := StringToUint("123", c)
	if result != uint(123) {
		t.Errorf("Expected: 123, Got: %d", result)
	}

	// Test invalid input
	result = StringToUint("invalid", c)
	if result != uint(0) {
		t.Errorf("Expected: 0, Got: %d", result)
	}
	// Add more test cases as needed
}

func TestFloat64ToInt(t *testing.T) {
	// Create a gin.Context for testing with gin.Default()
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	// Test case 1: Positive float
	result := Float64ToInt(123.45, c)
	if result != 123 {
		t.Errorf("Expected: 123, Got: %d", result)
	}

	// Test case 2: Negative float
	result = Float64ToInt(456.78, c)
	if result != 457 {
		t.Errorf("Expected: -456, Got: %d", result)
	}

	// Test case 3: Float with decimal part
	result = Float64ToInt(789.123, c)
	if result != 789 {
		t.Errorf("Expected: 789, Got: %d", result)
	}

	// Add more test cases as needed
}

func TestTimeToString(t *testing.T) {
	// Test case 1: Time with day < 10
	result := TimeToString(time.Now())
	if result != "1st Jan 2021 | 07:00 UTC" {
		t.Errorf("Expected: 1st Jan 2021 | 07:00 UTC, Got: %s", result)
	}

	// Test case 2: Time with day > 10
	result = TimeToString(time.Now())
	if result != "11th Jan 2021 | 07:00 UTC" {
		t.Errorf("Expected: 11th Jan 2021 | 07:00 UTC, Got: %s", result)
	}

	// Add more test cases as needed
}
