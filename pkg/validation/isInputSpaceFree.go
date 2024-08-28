package validation

import "strings"

func IsInputSpaceFree(input string) bool {
	if strings.Contains(input, " ") {

		return false
	}
	return true
}
