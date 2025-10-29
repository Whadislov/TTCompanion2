package my_functions

import (
	"log"

	mt "github.com/Whadislov/TTCompanion2/internal/my_types"
	"github.com/google/uuid"
)

// NewTeam creates a new team with the given name and adds it to the database.
// Returns the created team and an error if the team name is empty or if there is an issue with the operation.
func NewTeam(teamName string, db *mt.Database) (*mt.Team, error) {
	b, err := IsValidTeamClubName(teamName)
	if !b {
		return nil, err
	}

	b, err = IsStrTooLong(teamName, 30)
	if b {
		return nil, err
	}

	teamName = standardizeSpaces(teamName)

	t := &mt.Team{
		ID:        uuid.New(),
		Name:      teamName,
		PlayerIDs: make(map[uuid.UUID]string),
		ClubID:    make(map[uuid.UUID]string, 1), // Capacity 1
		IsNew:     true,
	}

	db.AddTeam(t)
	log.Printf("Team %v sucessfully created.", teamName)
	return t, nil
}
