package validation

import (
	"fmt"
	"strings"
)

// Function to check for single word username
func IsSingleWordUsername(input string) bool {
	if strings.Contains(input, " ") {
		fmt.Println("\033[1;31m\nInvalid Input\033[0m")
		fmt.Println("\nTry again....")
		return false
	}
	return true
}
