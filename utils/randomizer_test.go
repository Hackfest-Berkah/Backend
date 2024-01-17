package utils

import (
	"fmt"
	"testing"
)

func TestRandomUUIDString(t *testing.T) {
	result := RandomUUIDString()
	if result == "" {
		t.Errorf("Expected: not empty, Got: %s", result)
	}

	fmt.Println(result)
}
