package validate

import (
	"net/http"
	"regexp"
	"strconv"
)

var (
	nonNumbersRegex, _ = regexp.Compile("\\D")
)

func IsValidText(text string) bool {
	if isStringEmpty(text) || len(text) > 256 {
		return false
	}
	return true
}

func IsValidUri(uri string) bool {
	if isStringEmpty(uri) {
		return false
	}
	resp, err := http.Get(uri)
	if err != nil || resp.StatusCode != http.StatusOK {
		return false
	}
	return true
}

func IsValidNumber(stringToCheck string) bool {
	if isStringEmpty(stringToCheck) {
		return false
	}
	numberToCheck, err := strconv.Atoi(nonNumbersRegex.ReplaceAllString(stringToCheck, ""))
	if err != nil {
		return false
	}
	if numberToCheck < 0 {
		return false
	}
	return true
}

func isStringEmpty(s string) bool {
	if s == "" {
		return false
	}
	return true
}
