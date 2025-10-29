package my_functions

import (
	"fmt"

	mt "github.com/Whadislov/TTCompanion2/internal/my_types"
)

// RemovePlayerFromClub removes a player from a club by updating both the player's and the club's records.
// Returns an error if there is an issue with the operation.
func RemovePlayerFromClub(p *mt.Player, c *mt.Club) error {

	err := c.RemovePlayer(p)
	if err != nil {
		return fmt.Errorf("%s %s has not been successfully removed from %s. Reason : %w", p.Firstname, p.Lastname, c.Name, err)
	}

	err = p.RemoveClub(c)
	if err != nil {
		return fmt.Errorf("%s %s has not been successfully removed from %s. Reason : %w", p.Firstname, p.Lastname, c.Name, err)
	}
	return nil
}
