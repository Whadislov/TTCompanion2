package my_functions

import (
	"fmt"
	"log"
	"time"

	mt "github.com/Whadislov/TTCompanion2/internal/my_types"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// NewUser creates a new user with the given name and adds it to the database.
// Returns the created user and an error if the user name is empty or if there is an issue with the operation.
func NewUser(username string, email string, password string, confirmPassword string, db *mt.Database) (*mt.User, error) {
	// In the UI, Email is first asked, then username, then password
	b, err := IsValidEmail(email)
	if !b {
		return nil, err
	}
	log.Println("User creation : Email is valid.")

	b, err = IsValidUsername(username)
	if !b {
		return nil, err
	}
	b, err = IsStrTooLong(username, 30)
	if b {
		return nil, err
	}
	log.Println("User creation : Username is valid.")

	for _, user := range db.Users {
		if user.Name == username {
			return nil, fmt.Errorf("username is already taken")
		} else if user.Email == email {
			return nil, fmt.Errorf("email is already used")
		}
	}

	b, err = IsValidPassword(password)
	if !b {
		return nil, err
	}
	log.Println("User creation : Password is valid.")

	if password != confirmPassword {
		return nil, fmt.Errorf("passwords do not match")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	// ISO 8601 timestamp
	timestamp := time.Now().Format(time.RFC3339)

	u := &mt.User{
		ID:           uuid.New(),
		Name:         username,
		Email:        email,
		PasswordHash: string(hashedPassword),
		CreatedAt:    timestamp,
		IsNew:        true,
	}

	db.AddUser(u)
	log.Printf("User creation : User %v sucessfully created.", username)
	return u, nil
}
