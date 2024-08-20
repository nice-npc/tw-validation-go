package validator

import (
	"testing"
)

func TestValid(t *testing.T) {
	tests := []struct {
		name               string
		identityCardNumber string
		expected           bool
	}{
		{"ValidID_1", "A123456789", true},
		{"ValidID_2", "B234567895", true},
		{"InvalidID_WrongFormat", "Z123456789", false},
		{"InvalidID_WrongChecksum", "A123456788", false},
		{"EmptyID", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := Valid(tt.identityCardNumber)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestValidFormat(t *testing.T) {
	tests := []struct {
		name                 string
		taiwanIdentityNumber string
		expected             bool
	}{
		{"ValidFormat_1", "A123456789", true},
		{"ValidFormat_2", "B234567895", true},
		{"InvalidFormat_TooShort", "A12345678", false},
		{"InvalidFormat_NoLetter", "1234567890", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validFormat(tt.taiwanIdentityNumber)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestCalculateCheckSum(t *testing.T) {
	tests := []struct {
		name                string
		identityNumberChars []rune
		expected            int
	}{
		{"ValidChecksum_1", []rune("A123456789"), 9},
		{"ValidChecksum_2", []rune("B234567895"), 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateCheckSum(tt.identityNumberChars)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
