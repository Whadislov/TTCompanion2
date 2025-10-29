package my_functions

import (
	"log"

	mt "github.com/Whadislov/TTCompanion2/internal/my_types"
	"github.com/google/uuid"
)

// NewClub creates a new club with the given name and adds it to the database.
// Returns the created club and an error if the club name is empty or if there is an issue with the operation.
func NewClub(clubName string, db *mt.Database) (*mt.Club, error) {
	b, err := IsValidTeamClubName(clubName)
	if !b {
		return nil, err
	}

	b, err = IsStrTooLong(clubName, 30)
	if b {
		return nil, err
	}

	clubName = standardizeSpaces(clubName)

	c := &mt.Club{
		ID:        uuid.New(),
		Name:      clubName,
		PlayerIDs: make(map[uuid.UUID]string),
		TeamIDs:   make(map[uuid.UUID]string),
		IsNew:     true,
	}

	db.AddClub(c)
	log.Printf("Club %v sucessfully created.", clubName)
	return c, nil
}
