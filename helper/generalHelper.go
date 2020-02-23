package helper

import (
	"github.com/Mohannad-Zeido/HackerNewsScraper/types"
	"strconv"
)

func ExtractNumberFromString(s string) (int, bool) {

	number := types.NonNumbersRegex.ReplaceAllString(s, "")
	if number == "" {
		number = "-1"
	}
	n, err := strconv.Atoi(number)
	if err != nil {
		return 0, false
	}
	return n, true
}
