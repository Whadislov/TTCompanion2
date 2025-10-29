package my_functions

import (
	"fmt"
	"regexp"
	"strings"

	mt "github.com/Whadislov/TTCompanion2/internal/my_types"
)

// DefaultPlayerMaterial returns a slice of strings representing the default material for a player.
func DefaultPlayerMaterial() []string {
	defaultMaterial := ""
	return []string{defaultMaterial, defaultMaterial, defaultMaterial}
}

// GetName returns the name of the given entity (Player, Team, or Club).
// Returns an empty string if the entity type is not recognized.
func GetName(x interface{}) string {
	switch v := x.(type) {
	case mt.Player:
		{
			return v.Firstname + v.Lastname
		}
	case *mt.Player:
		{
			return v.Firstname + v.Lastname
		}
	case mt.Team:
		{
			return v.Name
		}
	case *mt.Team:
		{
			return v.Name
		}
	case mt.Club:
		{
			return v.Name
		}
	case *mt.Club:
		{
			return v.Name
		}
	default:
		{
			return ""
		}
	}
}

// isValidName verifies that the name follows some criterias
func IsValidName(name string) (bool, error) {
	if name == "" {
		return false, fmt.Errorf("name cannot be empty")
	}

	// Name can be composed (not mandatory with the ()), will start will a -
	// * means that the group can be repeated
	nameRegex := `^[a-zA-ZéèêçàÉÈÊÇÀßöäüÖÜÄ]+(-[a-zA-ZéèêçàÉÈÊÇÀßöäüÖÜÄ]+)*$`

	// Compile the regex
	re := regexp.MustCompile(nameRegex)

	// Verify if the string matches the regex
	if re.MatchString(name) {
		return re.MatchString(name), nil
	} else {
		return re.MatchString(name), fmt.Errorf("name must be valid. For exemple, Jean-François, Evelin")
	}
}

// IsValidUsername verifies that the name follows some criterias
func IsValidUsername(username string) (bool, error) {
	if username == "" {
		return false, fmt.Errorf("username cannot be empty")
	}

	usernameRegex := `^[a-zA-Z0-9_]+$`

	// Compile the regex
	re := regexp.MustCompile(usernameRegex)

	// Verify if the string matches the regex
	if re.MatchString(username) {
		return re.MatchString(username), nil
	} else {
		return re.MatchString(username), fmt.Errorf("username must be valid (only letters and figures are allowed, spaces are not allowed)")
	}
}

// isValidEmail verifies that the name follows a valid regex
func IsValidEmail(email string) (bool, error) {
	if email == "" {
		return false, fmt.Errorf("e-mail cannot be empty")
	}

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	// Compile the regex
	re := regexp.MustCompile(emailRegex)

	// Verify if the string matches the regex, true means yes
	if re.MatchString(email) {
		return re.MatchString(email), nil
	} else {
		return re.MatchString(email), fmt.Errorf("e-mail must be valid. Example: abc@def.com")
	}
}

// IsValidPassword verifies that the password is not empty and does no contain spaces
func IsValidPassword(password string) (bool, error) {
	if password == "" {
		return false, fmt.Errorf("password cannot be empty")
	}

	for _, char := range password {
		if char == ' ' {
			return false, fmt.Errorf("password must be valid (spaces are not allowed)")
		}
	}

	return true, nil
}

// IsValidTeamClubName verifies that the team name or club name is not empty and only contain letters and figures
func IsValidTeamClubName(name string) (bool, error) {
	if name == "" {
		return false, fmt.Errorf("name cannot be empty")
	}

	nameRegex := `^[a-zA-Z0-9 ]+$`

	// Compile the regex
	re := regexp.MustCompile(nameRegex)

	// Verify if the string matches the regex
	if re.MatchString(name) {
		return re.MatchString(name), nil
	} else {
		return re.MatchString(name), fmt.Errorf("name must be valid (letters, figures and one space are allowed)")
	}
}

// IsStrTooLong verifies that the string is not too long
func IsStrTooLong(s string, maxLength int) (bool, error) {
	if len(s) > maxLength {
		return true, fmt.Errorf("string is too long")
	}
	return false, nil
}

// standardizeSpaces removes spaces at the beginning and end of the string and replaces multiple spaces by one
func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
