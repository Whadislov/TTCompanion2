package my_types

import (
	"fmt"
	"github.com/google/uuid"
)

// SetPlayerID sets the ID of the player.
func (p *Player) SetPlayerID(id uuid.UUID) {
	p.ID = id
}

// SetPlayerName sets the first name of the player.
func (p *Player) SetPlayerFirstname(name string) {
	p.Firstname = name
}

// SetPlayerName sets the last name of the player.
func (p *Player) SetPlayerLastname(name string) {
	p.Lastname = name
}

// SetPlayerAge sets the age of the player.
func (p *Player) SetPlayerAge(age int) {
	p.Age = age
}

// SetPlayerRanking sets the ranking of the player.
func (p *Player) SetPlayerRanking(ranking int) {
	p.Ranking = ranking
}

// SetPlayerMaterial sets the material (forehand, backhand, blade) of the player.
func (p *Player) SetPlayerMaterial(forehand string, backhand string, blade string) {
	material := []string{forehand, backhand, blade}
	p.Material = material
}

// AddTeam adds a team to the player's list of teams.
// Returns an error if the team is already in the player's list.
func (p *Player) AddTeam(team *Team) error {
	if _, ok := p.TeamIDs[team.ID]; ok {
		return fmt.Errorf("player %v %v is already in team %v", p.Firstname, p.Lastname, team.Name)
	}

	if p.TeamIDs == nil {
		p.TeamIDs = make(map[uuid.UUID]string)
	}

	p.TeamIDs[team.ID] = team.Name
	return nil
}

// AddClub adds a club to the player's list of clubs.
// Returns an error if the club is already in the player's list.
func (p *Player) AddClub(club *Club) error {
	if _, ok := p.ClubIDs[club.ID]; ok {
		return fmt.Errorf("player %v %v is already in club %v", p.Firstname, p.Lastname, club.Name)
	}

	if p.ClubIDs == nil {
		p.ClubIDs = make(map[uuid.UUID]string)
	}

	p.ClubIDs[club.ID] = club.Name
	return nil
}

// RemoveTeam removes a team from the player's list of teams.
// Returns an error if the team is not in the player's list.
func (p *Player) RemoveTeam(team *Team) error {
	if _, ok := p.TeamIDs[team.ID]; !ok {
		return fmt.Errorf("player %v %v is not in team %v", p.Firstname, p.Lastname, team.Name)
	}
	delete(p.TeamIDs, team.ID)
	return nil
}

// RemoveClub removes a club from the player's list of clubs.
// Returns an error if the club is not in the player's list.
func (p *Player) RemoveClub(club *Club) error {
	if _, ok := p.ClubIDs[club.ID]; !ok {
		return fmt.Errorf("player %v %v is not in club %v", p.Firstname, p.Lastname, club.Name)
	}
	delete(p.ClubIDs, club.ID)
	return nil
}

// HasTeam returns True if the player has at least one team.
func (p *Player) HasTeam() bool {
	return len(p.TeamIDs) > 0

}

// HasClub returns True if the player has at least one club.
func (p *Player) HasClub() bool {
	return len(p.ClubIDs) > 0

}

// GetName returns the player's full name (first name + last name).
func (p *Player) GetName() string {
	return p.Firstname + p.Lastname
}
