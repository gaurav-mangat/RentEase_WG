package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"rentease/internal/domain/entities"
	"strconv"
	"strings"
)

// Creating an active user (only username)
var ActiveUser string

// Creating an active uesr object
var ActiveUserobject entities.User

// ReadInput reads input from the user with a prompt.
func ReadInput(prompt string) string {
	reader1 := bufio.NewReader(os.Stdin)

	// Display the prompt message
	fmt.Print(prompt)

	// Read the input until the newline character
	input, err := reader1.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}

	// Trim the newline character and any surrounding white spaces
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
func ReadPincode(n int) int {
	var pincode int
	var pincodeTemp string

	for {
		if n == 1 {
			pincodeTemp = ReadInput("Enter a 6-digit pincode for that area (enter to skip): ")
			if pincodeTemp == "" {
				pincodeTemp = "555555" // random pincode which is invalid
			}
		} else {
			pincodeTemp = ReadInput("Enter a 6-digit pincode (mandatory): ")
		}
		var err error
		pincode, err = strconv.Atoi(pincodeTemp)

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

// DisplayProperties prints the details of properties in a structured format.
func DisplayProperties(properties []entities.Property) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Property Type", "Title", "Address", "Rent Amount", "Approved By Admin", "Details"})

	table.SetAutoWrapText(true)

	// Set table border and separator styles
	table.SetBorder(true)         // Add border around the table
	table.SetHeaderLine(true)     // Add a line below the header
	table.SetColumnSeparator("|") // Set column separator
	table.SetCenterSeparator("|") // Set center separator
	table.SetRowSeparator("-")    // Set row separator
	table.SetRowLine(true)        // Add a line between rows

	// Add rows to the table
	for i, property := range properties {
		details := getPropertyDetails(property)
		address := formatAddress(property.Address)
		table.Append([]string{
			fmt.Sprintf("%d", i+1),
			propertyTypeToString(property.PropertyType),
			property.Title,
			address,
			fmt.Sprintf("%.2f", property.RentAmount),
			fmt.Sprintf("%t", property.IsApprovedByAdmin),
			details,
		})
	}

	table.Render()
}

func propertyTypeToString(propertyType int) string {
	switch propertyType {
	case 1:
		return "Commercial"
	case 2:
		return "House"
	case 3:
		return "Flat"
	default:
		return "Unknown"
	}
}

func formatAddress(address entities.Address) string {
	return fmt.Sprintf("%s, %s, %s, %d", address.Area, address.City, address.State, address.Pincode)
}

func getPropertyDetails(property entities.Property) string {
	switch details := property.Details.(type) {
	case entities.CommercialDetails:
		return fmt.Sprintf("Floor Area: %s, Subtype: %s", details.FloorArea, details.SubType)
	case entities.HouseDetails:
		return fmt.Sprintf("Rooms: %d, Furnished: %s, Amenities: %v", details.NoOfRooms, details.FurnishedCategory, details.Amenities)
	case entities.FlatDetails:
		return fmt.Sprintf("Furnished: %s, Amenities: %v, BHK: %d", details.FurnishedCategory, details.Amenities, details.BHK)
	default:
		return "Unknown details"
	}
}

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

func GetHiddenInput(prompt string) (string, error) {
	fmt.Print(prompt)
	bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Errorf("Error reading password: " + err.Error())
		return "", err
	}
	fmt.Println() // Print a newline after input
	return strings.TrimSpace(string(bytePassword)), nil
}

// ReadBHKInput takes BHK input between 1 - 6
func ReadBHKInput() (int, error) {
	var bhk int
	for {

		bhkTemp := ReadInput("Enter BHK (1-6): ")
		var err error
		bhk, err = strconv.Atoi(bhkTemp)
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

func GetAddressFromPincode(pincode int) (entities.Address, error) {
	url := fmt.Sprintf("https://api.postalpincode.in/pincode/%d", pincode)
	resp, err := http.Get(url)
	if err != nil {
		return entities.Address{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return entities.Address{}, err
	}

	var apiResponse []struct {
		Message    string `json:"Message"`
		Status     string `json:"Status"`
		PostOffice []struct {
			Name     string `json:"Name"`
			District string `json:"District"`
			State    string `json:"State"`
			Pincode  string `json:"Pincode"`
		} `json:"PostOffice"`
	}

	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return entities.Address{}, err
	}

	if len(apiResponse) == 0 || len(apiResponse[0].PostOffice) == 0 {
		return entities.Address{}, fmt.Errorf("no address details found for pincode %d", pincode)
	}

	postOffice := apiResponse[0].PostOffice[0]
	return entities.Address{
		Area:    postOffice.Name,
		City:    postOffice.District,
		State:   postOffice.State,
		Pincode: pincode,
	}, nil
}
