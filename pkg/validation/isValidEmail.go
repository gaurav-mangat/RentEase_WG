package validation

import "regexp"

// IsValidEmail validates the email format using a regular expression
func IsValidEmail(email string) bool {
	// Basic email validation regex
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
