package my_types

import (
	"fmt"
	"log"
)

func (p *Player) String() string {
	return fmt.Sprintf("Id: %v, Firstname: %s, Lastname: %s, Age: %d, Ranking: %d, Material: %v, Teams: %v, Clubs: %v",
		p.ID,
		p.Firstname,
		p.Lastname,
		p.Age,
		p.Ranking,
		p.Material,
		p.TeamIDs,
		p.ClubIDs,
	)
}

func (p *Player) Show() error {
	if p == nil {
		log.Println("Error! Player does not exist (nil pointer)")
		return fmt.Errorf("player does not exist")
	}

	fmt.Println("Showing characteristics of", p.Firstname, p.Lastname)

	// Get team name
	teams := []string{}
	for _, team := range p.TeamIDs {
		teams = append(teams, team)
	}

	// Get club name
	clubs := []string{}
	for _, club := range p.ClubIDs {
		clubs = append(clubs, club)
	}

	fmt.Printf("%s %s, Age: %d, Ranking: %d, Material: %v, Teams: %v, Clubs: %v.\n",
		p.Firstname,
		p.Lastname,
		p.Age,
		p.Ranking,
		p.Material,
		teams,
		clubs,
	)
	return nil
}

func (t *Team) Show() error {
	if t == nil {
		log.Println("Error! Team does not exist (nil pointer)")
		return fmt.Errorf("team does not exist")
	}

	fmt.Println("Showing the characteristics of", t.Name)
	n := len(t.PlayerIDs)
	if n >= 1 {
		for i, player := range t.PlayerIDs {
			fmt.Printf("Player %v: %v\n",
				i,
				player,
			)
		}
	} else {
		fmt.Printf("There is no player in %v.\n",
			t.Name,
		)
	}
	return nil
}

func (c *Club) Show() error {
	if c == nil {
		log.Println("Error! Club does not exist (nil pointer)")
		return fmt.Errorf("club does not exist")
	}

	fmt.Println("Characteristics of", c.Name)
	n := len(c.TeamIDs)
	if n >= 1 {
		for i, team := range c.TeamIDs {
			fmt.Printf("Team %v: %v.\n",
				i,
				team,
			)
		}
	} else {
		fmt.Printf("There is no team in %v.\n",
			c.Name,
		)
	}
	return nil
}

func (d *Database) Show() error {
	fmt.Println("Clubs :")
	for _, club := range d.Clubs {
		err1 := club.Show()
		if err1 != nil {
			return err1
		}
	}
	fmt.Println("Teams :")
	for _, team := range d.Teams {
		err2 := team.Show()
		if err2 != nil {
			return err2
		}
	}
	fmt.Println("Players :")
	for _, player := range d.Players {
		err3 := player.Show()
		if err3 != nil {
			return err3
		}
	}
	return nil
}
