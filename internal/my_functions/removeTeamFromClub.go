package my_functions

import (
	"fmt"

	mt "github.com/Whadislov/TTCompanion2/internal/my_types"
)

// RemoveTeamFromClub removes a team from a club by updating both the team's and the club's records.
// Returns an error if there is an issue with the operation.
func RemoveTeamFromClub(t *mt.Team, c *mt.Club) error {

	err := c.RemoveTeam(t)
	if err != nil {
		return fmt.Errorf("%s has not been successfully removed from %s. Reason : %w", t.Name, c.Name, err)
	}

	err = t.RemoveClub(c)
	if err != nil {
		return fmt.Errorf("%s has not been successfully removed from %s. Reason : %w", t.Name, c.Name, err)
	}
	return nil
}
