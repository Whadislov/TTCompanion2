package my_functions

import (
	"fmt"

	mt "github.com/Whadislov/TTCompanion2/internal/my_types"
)

// AddTeamToClub adds a team to a club by updating both the team's and the club's records.
// Returns an error if the team is already in the club or if there is an issue with the operation.
func AddTeamToClub(t *mt.Team, c *mt.Club) error {

	err := t.AddClub(c)
	if err != nil {
		return fmt.Errorf("error when adding team %v to club %v: %w", t.Name, c.Name, err)
	}
	err = c.AddTeam(t)
	if err != nil {
		return fmt.Errorf("error when adding team %v to club %v: %w", t.Name, c.Name, err)
	}

	return nil
}
