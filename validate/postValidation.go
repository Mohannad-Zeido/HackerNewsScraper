package validate

import (
	"fmt"
	"net/url"
)

func IsValidText(text string) bool {
	if text == "" || len(text) > 256 {
		fmt.Println("Title '" + text + "' failed validation")
		return false
	}
	return true
}

func IsValidUri(uri string) bool {
	_, err := url.ParseRequestURI(uri)
	if err != nil {
		fmt.Println("uri '" + uri + "' failed validation")
		return false
	}
	return true

}

func IsValidNumber(numberToCheck int) bool {
	if numberToCheck < 0 {
		fmt.Println("number '" + string(numberToCheck) + "' failed validation")
		return false
	}
	return true
}
