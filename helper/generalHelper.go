package helper

import (
	"github.com/Mohannad-Zeido/HackerNewsScraper/types"
	"regexp"
	"strconv"
)

//the compiled regular expression
var nonNumbersRegex, _ = regexp.Compile(types.NonNumbers)

//ExtractNumberFromString will return the number that is present in the string.
//an empty string is the indication that a number is not present. as a result the return value will be -1
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
