package validate

import (
	"net/url"
)

func IsValidText(text string) bool {
	if text == "" || len(text) > 256 {
		return false
	}
	return true
}

func IsValidUri(uri string) bool {
	_, err := url.ParseRequestURI(uri)
	if err != nil {
		return false
	}
	return true

}

func IsValidNumber(numberToCheck int) bool {
	if numberToCheck < 0 {
		return false
	}
	return true
}
