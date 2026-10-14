package pkg

import (
	"os"
	"strings"
)

// GetAllowedMobileNumbers returns a slice of directory names in test_data_dir
func GetAllowedMobileNumbers() []string {
	dirEntries, err := os.ReadDir("test_data_dir")
	if err != nil {
		return nil
	}
	var numbers []string
	for _, entry := range dirEntries {
		if entry.IsDir() {
			numbers = append(numbers, entry.Name())
		}
	}
	return numbers
}

// MobileNumberAllowed reports whether phone is one of the dummy personas.
// The login page checks this in the browser. The login handler must check it
// too, because a direct POST can skip that script.
func MobileNumberAllowed(phone string, allowed []string) bool {
	phone = strings.TrimSpace(phone)
	if phone == "" {
		return false
	}
	for _, candidate := range allowed {
		if candidate == phone {
			return true
		}
	}
	return false
}
