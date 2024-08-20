package validator

import (
	"errors"
	"regexp"
	"strings"
)

const LAST_ID_INDEX = 9

var REGEX = "[A-Z][12]\\d{8}"
var PATTERN = regexp.MustCompile(REGEX)
var MULTIPLIERS = []int{8, 7, 6, 5, 4, 3, 2, 1}

var FIRST_LETTER_VALUE_MAP = map[string]int{
	"A": 10, "B": 11, "C": 12, "D": 13, "E": 14,
	"F": 15, "G": 16, "H": 17, "I": 34, "J": 18,
	"K": 19, "L": 20, "M": 21, "N": 22, "O": 35,
	"P": 23, "Q": 24, "R": 25, "S": 26, "T": 27,
	"U": 28, "V": 29, "W": 32, "X": 30, "Y": 31,
	"Z": 33,
}

// Valid validates the identity card number
func Valid(identityCardNumber string) (bool, error) {
	if identityCardNumber == "" {
		return false, errors.New("identity card number is empty")
	}

	upperCaseIdentityCardNumber := strings.ToUpper(identityCardNumber)

	if !validFormat(upperCaseIdentityCardNumber) {
		return false, nil
	}

	identityNumberChars := []rune(upperCaseIdentityCardNumber)
	checkSum := int(identityNumberChars[LAST_ID_INDEX] - '0')
	calculateCheckSum := calculateCheckSum(identityNumberChars)
	return calculateCheckSum == checkSum, nil
}

// validFormat checks the format of the identity card number
func validFormat(taiwanIdentityNumber string) bool {
	return PATTERN.MatchString(taiwanIdentityNumber)
}

// calculateCheckSum calculates the checksum of the identity card number
func calculateCheckSum(identityNumberChars []rune) int {
	firstLetter := string(identityNumberChars[0])
	firstLetterValue := FIRST_LETTER_VALUE_MAP[firstLetter]

	sum := (firstLetterValue / 10) + (firstLetterValue%10)*9

	for i := 1; i < 9; i++ {
		sum += int(identityNumberChars[i]-'0') * MULTIPLIERS[i-1]
	}

	sumMod10 := sum % 10
	if sumMod10 == 0 {
		return 0
	}
	return 10 - sumMod10
}
