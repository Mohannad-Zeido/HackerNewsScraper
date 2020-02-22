package validate

import "net/http"

func IsValidText(text string) bool {
	if text == "" || len(text) < 256 {
		return false
	}
	return true
}

func IsValidUri(uri string) bool {
	resp, err := http.Get(uri)
	if err != nil || resp.StatusCode != http.StatusOK {
		return false
	}
	return true
}

func IsValidNumber(number int) bool {
	if number < 0 {
		return false
	}
	return true
}
