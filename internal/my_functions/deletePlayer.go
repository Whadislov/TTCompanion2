package my_functions

import (
	"fmt"

	mt "github.com/Whadislov/TTCompanion2/internal/my_types"
	"github.com/google/uuid"
)

// DeletePlayer removes a player from the database and updates all related team and club records.
// Returns an error if there is an issue with the operation.
func DeletePlayer(p *mt.Player, db *mt.Database) error {
	// Remove club depedences
	var clubIDs []uuid.UUID
	if len(p.ClubIDs) > 0 {
		for clubID := range p.ClubIDs {
			clubIDs = append(clubIDs, clubID)
			err := db.Clubs[clubID].RemovePlayer(p)
			if err != nil {
				return fmt.Errorf("error when deleting player %s %s: %w", p.Firstname, p.Lastname, err)
			}
		}
		for _, clubID := range clubIDs {
			err := p.RemoveClub(db.Clubs[clubID])
			if err != nil {
				return fmt.Errorf("error when deleting player %s %s: %w", p.Firstname, p.Lastname, err)
			}
		}
	}

	// Remove team depedences
	var teamIDs []uuid.UUID
	if len(p.TeamIDs) > 0 {
		for teamID := range p.TeamIDs {
			teamIDs = append(teamIDs, teamID)
			err := db.Teams[teamID].RemovePlayer(p)
			if err != nil {
				return fmt.Errorf("error when deleting player %s %s: %w", p.Firstname, p.Lastname, err)
			}
		}
		for _, teamID := range teamIDs {
			err := p.RemoveTeam(db.Teams[teamID])
			if err != nil {
				return fmt.Errorf("error when deleting player %s %s: %w", p.Firstname, p.Lastname, err)
			}
		}
	}

	// Delete player
	IDtoDelete := p.ID
	// Delete from the local database
	err := db.DeletePlayer(p.ID)
	if err != nil {
		return fmt.Errorf("error when deleting player %s %s: %w", p.Firstname, p.Lastname, err)
	} else {
		// If already in postgres, store the ID to be deleted
		if !p.IsNew {
			db.AddDeletedPlayer(IDtoDelete)
		}
	}

	return nil
}
