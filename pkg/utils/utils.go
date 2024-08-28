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

// Creating an active user (only username)
var ActiveUser string

// Creating an active uesr object
var ActiveUserobject entities.User

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

// ReadPincode read the pincode and makes it mandatory to enter pincode
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

	fmt.Print("\nProperty Type:  ")
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
	fmt.Printf("Expected Rent Amount: %.2f\n", property.RentAmount)
	fmt.Println("Is Property approved : ", property.IsApprovedByAdmin)

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
		fmt.Println(details)
		//fmt.Println("  Unknown property details")
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

// ReadBHKInput takes BHK input between 1 - 6
func ReadBHKInput() (int, error) {
	var bhk int
	for {
		fmt.Print("Enter BHK (1-6): ")
		_, err := fmt.Scanf("%d", &bhk)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}
		if bhk < 1 || bhk > 6 {
			fmt.Println("BHK must be between 1 and 6.")
			continue
		}
		return bhk, nil
	}
}
