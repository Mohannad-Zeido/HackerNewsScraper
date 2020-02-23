package validate

import (
	"net/url"
)

// IsValidText will return false if the text is empty or the length is greater than 256
func IsValidText(text string) bool {
	if text == "" || len(text) > 256 {
		return false
	}

	return true
}

//IsValidURI checks the uri is valid according to the  RFC 3986 standard
func IsValidURI(uri string) bool {
	_, err := url.ParseRequestURI(uri)
	return !(err != nil)
}

//IsValidNumber will return false if the number is less than 0
func IsValidNumber(numberToCheck int) bool {
	return !(numberToCheck < 0)
}
