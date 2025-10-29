package my_functions

import (
	"fmt"

	mt "github.com/Whadislov/TTCompanion2/internal/my_types"
	"github.com/google/uuid"
)

// DeleteTeam removes a team from the database and updates all related player and club records.
// Returns an error if there is an issue with the operation.
func DeleteTeam(t *mt.Team, db *mt.Database) error {
	// Remove club depedences
	if len(t.ClubID) > 0 {
		// need clubID for t.RemoveClub
		var clubID uuid.UUID
		for ID := range t.ClubID {
			clubID = ID
			err := db.Clubs[ID].RemoveTeam(t)
			if err != nil {
				return fmt.Errorf("error when deleting team %s: %w", t.Name, err)
			}
		}
		err := t.RemoveClub(db.Clubs[clubID])
		if err != nil {
			return fmt.Errorf("error when deleting team %s: %w", t.Name, err)
		}
	}

	// Remove player depedences
	var playerIDs []uuid.UUID
	if len(t.PlayerIDs) > 0 {
		for playerID := range t.PlayerIDs {
			playerIDs = append(playerIDs, playerID)
			err := db.Players[playerID].RemoveTeam(t)
			if err != nil {

				return fmt.Errorf("error when deleting team %s: %w", t.Name, err)
			}
		}
		for _, playerID := range playerIDs {
			err := t.RemovePlayer(db.Players[playerID])
			if err != nil {
				return fmt.Errorf("error when deleting team %s: %w", t.Name, err)
			}
		}
	}

	// Delete team
	IDtoDelete := t.ID
	// Delete from the local database
	err := db.DeleteTeam(t.ID)
	if err != nil {
		return fmt.Errorf("error when deleting team %s: %w", t.Name, err)
	} else {
		// If already in postgres, store the ID to be deleted
		if !t.IsNew {
			db.AddDeletedTeam(IDtoDelete)
		}
	}

	return nil
}
