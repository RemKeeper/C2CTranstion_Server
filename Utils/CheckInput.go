package Utils

import (
	"regexp"
)

func IsValidInput(input string) bool {
	pattern := "^[a-zA-Z0-9_]*$"
	matched, _ := regexp.MatchString(pattern, input)
	return matched
}
