package my_functions

import (
	"fmt"

	mt "github.com/Whadislov/TTCompanion2/internal/my_types"
)

// AddPlayerToClub adds a player to a club by updating both the player's and the club's records.
// Returns an error if the player is already in the club or if there is an issue with the operation.
func AddPlayerToClub(p *mt.Player, c *mt.Club) error {

	err := p.AddClub(c)
	if err != nil {
		return fmt.Errorf("error when adding player %v %v to club %v: %w", p.Firstname, p.Lastname, c.Name, err)
	}
	err = c.AddPlayer(p)
	if err != nil {
		return fmt.Errorf("error when adding player %v %v to club %v: %w", p.Firstname, p.Lastname, c.Name, err)
	}

	return nil
}
