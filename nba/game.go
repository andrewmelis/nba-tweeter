package nba

import (
	"fmt"

	"github.com/andrewmelis/nba-tweeter/play"
)

type NBAGames struct {
	Games []*NBAGame `json:"games"`
}

type NBAGame struct {
	Visitor NBATeam `json:"vTeam"`
	Home    NBATeam `json:"hTeam"`
}

func (g *NBAGame) GameCode() string {
	return fmt.Sprintf("%s%s", g.Visitor.TriCode, g.Home.TriCode)
}

func (g *NBAGame) Plays() []play.Play {
	plays := make([]play.Play, 0)
	return plays
}

type NBATeam struct {
	TriCode string `json:"triCode"`
}
