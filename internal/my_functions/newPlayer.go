package my_functions

import (
	"log"
	"strings"
	"unicode"
	"unicode/utf8"

	mt "github.com/Whadislov/TTCompanion2/internal/my_types"
	"github.com/google/uuid"
)

// NewPlayer creates a new player with the given name and adds it to the database.
// Returns the created player and an error if the player name is empty or if there is an issue with the operation.
func NewPlayer(firstname string, lastname string, db *mt.Database) (*mt.Player, error) {
	b, err := IsValidName(firstname)
	if !b {
		return nil, err
	}

	b, err = IsStrTooLong(firstname, 40)
	if b {
		return nil, err
	}

	b, err = IsValidName(lastname)
	if !b {
		return nil, err
	}

	b, err = IsStrTooLong(lastname, 40)
	if b {
		return nil, err
	}

	firstname = strings.ToLower(firstname)
	firstRune, size := utf8.DecodeRuneInString(firstname)
	if firstRune != utf8.RuneError {
		firstname = string(unicode.ToUpper(firstRune)) + firstname[size:]
	}

	lastname = strings.ToLower(lastname)
	firstRune, size = utf8.DecodeRuneInString(lastname)
	if firstRune != utf8.RuneError {
		lastname = string(unicode.ToUpper(firstRune)) + lastname[size:]
	}

	p := &mt.Player{
		ID:        uuid.New(),
		Firstname: firstname,
		Lastname:  lastname,
		Age:       -1,
		Ranking:   -1,
		Material:  DefaultPlayerMaterial(),
		TeamIDs:   make(map[uuid.UUID]string),
		ClubIDs:   make(map[uuid.UUID]string),
		IsNew:     true,
	}

	db.AddPlayer(p)
	log.Printf("Player %v %v sucessfully created.", firstname, lastname)
	return p, nil
}
