package utils

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
	"os"
	"regexp"
	"rentease/internal/domain/entities"
	"strings"
)

// Creating an active user
var ActiveUser string

// Create a buffered reader
var Reader *bufio.Reader

// Initialize the Reader in the init function
func init() {
	Reader = bufio.NewReader(os.Stdin)
}

// ReadInput reads input from the user with a prompt.
func ReadInput(prompt string) string {
	fmt.Print(prompt)
	input, err := Reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}
	return strings.TrimSpace(input)
}

// Function to check for single word username
func IsValidInput2(input string) bool {
	if strings.Contains(input, " ") {
		fmt.Println("\033[1;31m\nInvalid Input\033[0m")
		fmt.Println("\nTry again....")
		return false
	}
	return true
}

func IsValidInput(input string) bool {
	if strings.Contains(input, " ") {

		return false
	}
	return true
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compares a password with its hash.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func IsValidPassword(password string) bool {
	var (
		hasUpper   = regexp.MustCompile(`[A-Z]`).MatchString
		hasLower   = regexp.MustCompile(`[a-z]`).MatchString
		hasNumber  = regexp.MustCompile(`[0-9]`).MatchString
		hasSpecial = regexp.MustCompile(`[!@#\$%\^&\*\(\)_+\-=\[\]\;:'",.<>?/|\\]`).MatchString
	)

	return len(password) > 8 && hasUpper(password) && hasLower(password) && hasNumber(password) && hasSpecial(password)
}

func IsValidMobileNumber(number string) bool {
	match, _ := regexp.MatchString(`^[6-9]\d{9}$`, number)
	return match
}

// isValidEmail validates the email format using a regular expression
func IsValidEmail(email string) bool {
	// Basic email validation regex
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

// Reading pincode
func ReadPincode() int {
	var pincode int

	for {
		fmt.Print("Enter a 6-digit pincode (mandatory): ")
		_, err := fmt.Scan(&pincode)

		// Check if there's an error in scanning or if the pincode is not 6 digits
		if err != nil || pincode < 100000 || pincode > 999999 {
			fmt.Println("Invalid pincode. Please enter a valid 6-digit pincode.")
			continue
		}

		// If valid, break the loop
		break
	}

	return pincode
}

// DisplayProperty prints the details of a single property in a structured format.
func DisplayProperty(property entities.Property) {

	fmt.Print("Property Type:  ")
	switch property.PropertyType {
	case 1:
		fmt.Println("Commercial")
	case 2:
		fmt.Println("House")
	case 3:
		fmt.Println("Flat")
	default:
		fmt.Println("Unknown")
	}

	fmt.Printf("Property Title: %s\n", property.Title)
	fmt.Printf("Address: %s, %s, %s, %d\n", property.Address.Area, property.Address.City, property.Address.State, property.Address.Pincode)
	fmt.Printf("Is Approved by Admin : %v\n", property.IsApprovedByAdmin)
	fmt.Printf("Expected Rent Amount: %f\n", property.RentAmount)

	fmt.Println("Other Details:")

	switch details := property.Details.(type) {
	case entities.CommercialDetails:
		fmt.Printf("  Floor Area: %s\n", details.FloorArea)
		fmt.Printf("  Subtype: %s\n", details.SubType)
	case entities.HouseDetails:
		fmt.Printf("  Number of Rooms: %d\n", details.NoOfRooms)
		fmt.Printf("  Furnished Category: %s\n", details.FurnishedCategory)
		fmt.Printf("  Amenities: %v\n", details.Amenities)
	case entities.FlatDetails:
		fmt.Printf("  Furnished Category: %s\n", details.FurnishedCategory)
		fmt.Printf("  Amenities: %v\n", details.Amenities)
		fmt.Printf("  BHK: %d\n", details.BHK)
	default:
		fmt.Println("  Unknown property details")
	}
	fmt.Println()

}

// DisplayProperties prints a list of properties using DisplayProperty for each.
func DisplayProperties(properties []entities.Property) {
	for i, property := range properties {
		fmt.Printf("Property #%d:\n", i+1)
		DisplayProperty(property)
	}
}

func DisplayPropertyshortInfo(properties []entities.Property) {
	// Display the list of properties with short info
	for i, property := range properties {
		fmt.Printf("Property #%d:\n", i+1)
		fmt.Printf("  Title: %s\n", property.Title)
		fmt.Printf("  Rent Amount: %f\n", property.RentAmount)
		fmt.Printf("  Address: %s, %s, %s, %d\n", property.Address.Area, property.Address.City, property.Address.State, property.Address.Pincode)
		fmt.Println()
	}

	//Continuously prompt the user to select a property number for more details until valid input is given
}

func GetHiddenInput(prompt string) string {
	fmt.Print(prompt)
	bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Errorf("Error reading password: " + err.Error())
		return ""
	}
	fmt.Println() // Print a newline after input
	return strings.TrimSpace(string(bytePassword))
}
