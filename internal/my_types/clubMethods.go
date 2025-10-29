package my_types

import (
	"fmt"
	"github.com/google/uuid"
)

// SetClubID sets the ID of the club.
func (c *Club) SetClubID(id uuid.UUID) {
	c.ID = id
}

// SetClubName sets the name of the club.
func (c *Club) SetClubName(name string) {
	c.Name = name
}

// AddPlayer adds a player to the club.
// Returns an error if the player is already in the club.
func (c *Club) AddPlayer(player *Player) error {
	if _, ok := c.PlayerIDs[player.ID]; ok {
		return fmt.Errorf("player %v is already in club %v", player.Firstname+player.Lastname, c.Name)
	}

	if c.PlayerIDs == nil {
		c.PlayerIDs = make(map[uuid.UUID]string)
	}

	c.PlayerIDs[player.ID] = player.Firstname + player.Lastname
	return nil
}

// AddTeam adds a team to the club.
// Returns an error if the team is already in the club.
func (c *Club) AddTeam(team *Team) error {
	if _, ok := c.TeamIDs[team.ID]; ok {
		return fmt.Errorf("team %v is already in club %v", team.Name, c.Name)
	}

	if c.TeamIDs == nil {
		c.TeamIDs = make(map[uuid.UUID]string)
	}

	c.TeamIDs[team.ID] = team.Name
	return nil
}

// RemovePlayer removes a player from the club.
// Returns an error if the player is not in the club.
func (c *Club) RemovePlayer(player *Player) error {
	if _, ok := c.PlayerIDs[player.ID]; !ok {
		return fmt.Errorf("player %v is not in club %v", player.Firstname+player.Lastname, c.Name)
	}
	delete(c.PlayerIDs, player.ID)
	return nil
}

// RemoveTeam removes a team from the club.
// Returns an error if the team is not in the club.
func (c *Club) RemoveTeam(team *Team) error {
	if _, ok := c.TeamIDs[team.ID]; !ok {
		return fmt.Errorf("team %v is not in club %v", team.Name, c.Name)
	}
	delete(c.TeamIDs, team.ID)
	return nil
}

// HasPlayer returns True if the club has at least one player.
func (c *Club) HasPlayer() bool {
	return len(c.PlayerIDs) > 0
}

// HasTeam returns True if the club has at least one team.
func (c *Club) HasTeam() bool {
	return len(c.TeamIDs) > 0
}

// GetName returns the club's name.
func (c *Club) GetName() string {
	return c.Name
}
