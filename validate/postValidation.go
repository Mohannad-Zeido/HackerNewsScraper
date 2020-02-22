package validate

import (
	"fmt"
	"net/http"
)

func IsValidText(text string) bool {
	if text == "" || len(text) > 256 {
		return false
	}
	return true
}

func IsValidUri(uri string) bool {

	fmt.Println("Testing Url: " + uri)
	resp, err := http.Get(uri)
	if err != nil || resp.StatusCode != http.StatusOK {
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
