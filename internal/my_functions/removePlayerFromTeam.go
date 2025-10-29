package my_functions

import (
	"fmt"

	mt "github.com/Whadislov/TTCompanion2/internal/my_types"
)

// RemovePlayerFromTeam removes a player from a team by updating both the player's and the team's records.
// Returns an error if there is an issue with the operation.
func RemovePlayerFromTeam(p *mt.Player, t *mt.Team) error {

	err := t.RemovePlayer(p)
	if err != nil {
		return fmt.Errorf("%s %s has not been successfully removed from %s. Reason : %w", p.Firstname, p.Lastname, t.Name, err)
	}

	err = p.RemoveTeam(t)
	if err != nil {
		return fmt.Errorf("%s %s has not been successfully removed from %s. Reason : %w", p.Firstname, p.Lastname, t.Name, err)
	}
	return nil
}
