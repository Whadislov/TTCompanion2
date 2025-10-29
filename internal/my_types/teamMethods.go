package my_types

import (
	"fmt"
	"github.com/google/uuid"
)

// SetTeamID sets the ID of the team.
func (t *Team) SetTeamID(id uuid.UUID) {
	t.ID = id
}

// SetTeamName sets the name of the team.
func (t *Team) SetTeamName(name string) {
	t.Name = name
}

// AddPlayer adds a player to the team.
// Returns an error if the player is already in the team.
func (t *Team) AddPlayer(player *Player) error {
	if _, ok := t.PlayerIDs[player.ID]; ok {
		return fmt.Errorf("player %v is already in team %v", player.Firstname+player.Lastname, t.Name)
	}

	if t.PlayerIDs == nil {
		t.PlayerIDs = make(map[uuid.UUID]string)
	}

	t.PlayerIDs[player.ID] = player.Firstname + player.Lastname
	return nil
}

// AddClub adds a club to the team.
// Returns an error if the team is already in a club.
func (t *Team) AddClub(club *Club) error {
	if len(t.ClubID) > 0 {
		return fmt.Errorf("team %v is already in a club", t.Name)
	}

	if t.ClubID == nil {
		t.ClubID = make(map[uuid.UUID]string)
	}

	t.ClubID[club.ID] = club.Name
	return nil
}

// RemovePlayer removes a player from the team.
// Returns an error if the player is not in the team.
func (t *Team) RemovePlayer(player *Player) error {
	if _, ok := t.PlayerIDs[player.ID]; !ok {
		return fmt.Errorf("player %v is not in team %v", player.Firstname+player.Lastname, t.Name)
	}
	delete(t.PlayerIDs, player.ID)
	return nil
}

// RemoveClub removes the club from the team.
// Returns an error if the team is not in a club.
func (t *Team) RemoveClub(club *Club) error {
	if _, ok := t.ClubID[club.ID]; !ok {
		return fmt.Errorf("team %v is not in club %v", t.Name, club.Name)
	}
	delete(t.ClubID, club.ID)
	return nil
}

// HasTeam returns True if the team has at least one player.
func (t *Team) HasPlayer() bool {
	return len(t.PlayerIDs) > 0
}

// HasClub returns True if the team has at least one club.
func (t *Team) HasClub() bool {
	return len(t.ClubID) > 0
}

// GetName returns the team's name.
func (t *Team) GetName() string {
	return t.Name
}
