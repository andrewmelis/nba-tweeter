package nba

import (
	"fmt"
)

type NBAGames struct {
	Games []NBAGame `json:"games"`
}

type NBAGame struct {
	Visitor NBATeam `json:"vTeam"`
	Home    NBATeam `json:"hTeam"`
}

func (g NBAGame) GameCode() string {
	return fmt.Sprintf("%s%s", g.Visitor.TriCode, g.Home.TriCode)

}

type NBATeam struct {
	TriCode string `json:"triCode"`
}
