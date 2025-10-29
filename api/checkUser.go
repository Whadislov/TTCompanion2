package api

import (
	"fmt"
	"log"

	mdb "github.com/Whadislov/TTCompanion2/internal/my_db"
	mt "github.com/Whadislov/TTCompanion2/internal/my_types"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// checkUserCredentials verifies if the credentials of the user are correct, returns the userID
func checkUserCredentials(username string, password string) (uuid.UUID, error) {
	// return a wrong ID if faulty
	var errorID uuid.UUID = [16]byte{0}
	log.Println("Loading DB to check user credentials")
	db, err := mdb.LoadUsersOnly()
	if err != nil {
		return errorID, fmt.Errorf("error during connexion to database to check user credentials")
	}

	for _, user := range db.Users {
		if username == user.Name {

			err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
			if err != nil {
				// Password is wrong
				log.Println("User credentials are wrong")
				return errorID, fmt.Errorf("username and password missmatch")
			} else {
				log.Println("User credentials are good")
				mdb.SetUserIDOfSession(user.ID)
				return user.ID, nil
			}
		}
	}
	// User does not exist
	return errorID, fmt.Errorf("username and password missmatch")
}

// checkUserCredentials verifies that the newly created user does not use an already registered username
func checkUserExists(username string, email string, db *mt.Database) (int, error) {
	log.Println("Loading DB to check user existence before creation")

	// Check Email first as email is the first field in the UI
	for _, user := range db.Users {
		if user.Email == email {
			return 1, fmt.Errorf("email already exists")
		}
		if user.Name == username {
			return 2, fmt.Errorf("username already exists")
		}

	}
	return 0, nil
}
