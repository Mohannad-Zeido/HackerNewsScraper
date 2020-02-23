package helper

import (
	"github.com/Mohannad-Zeido/HackerNewsScraper/types"
	"regexp"
	"strconv"
)

var nonNumbersRegex, _ = regexp.Compile(types.NonNumbers)

func ExtractNumberFromString(s string) (int, bool) {
	number := nonNumbersRegex.ReplaceAllString(s, "")
	if number == "" {
		number = "-1"
	}
	n, err := strconv.Atoi(number)
	if err != nil {
		return 0, false
	}
	return n, true
}
