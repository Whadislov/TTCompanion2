package my_functions

import (
	"fmt"

	mt "github.com/Whadislov/TTCompanion2/internal/my_types"
	"github.com/google/uuid"
)

// DeleteClub removes a club from the database and updates all related player and team records.
// Returns an error if there is an issue with the operation.
func DeleteClub(c *mt.Club, db *mt.Database) error {
	// Remove player depedences
	var playerIDs []uuid.UUID
	if len(c.PlayerIDs) > 0 {
		for playerID := range c.PlayerIDs {
			playerIDs = append(playerIDs, playerID)
			err := db.Players[playerID].RemoveClub(c)
			if err != nil {
				return fmt.Errorf("error when deleting club %s: %w", c.Name, err)
			}
		}
		for _, playerID := range playerIDs {
			err := c.RemovePlayer(db.Players[playerID])
			if err != nil {
				return fmt.Errorf("error when deleting club %s: %w", c.Name, err)
			}
		}
	}

	// Remove team depedences
	var teamIDs []uuid.UUID
	if len(c.TeamIDs) > 0 {
		for teamID := range c.TeamIDs {
			teamIDs = append(teamIDs, teamID)
			err := db.Teams[teamID].RemoveClub(c)
			if err != nil {
				return fmt.Errorf("error when deleting club %s: %w", c.Name, err)
			}
		}
		for _, teamID := range teamIDs {
			err := c.RemoveTeam(db.Teams[teamID])
			if err != nil {
				return fmt.Errorf("error when deleting club %s: %w", c.Name, err)
			}
		}
	}

	// Delete club
	IDtoDelete := c.ID
	// Delete from the local database
	err := db.DeleteClub(c.ID)
	if err != nil {
		return fmt.Errorf("error when deleting club %s: %w", c.Name, err)
	} else {
		// If already in postgres, store the ID to be deleted
		if !c.IsNew {
			db.AddDeletedClub(IDtoDelete)
		}
	}

	return nil
}
