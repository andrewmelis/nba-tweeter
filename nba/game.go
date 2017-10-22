package nba

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/andrewmelis/nba-tweeter/play"
)

var BaseURL = "https://data.nba.net"

type NBAGames struct {
	Games []*NBAGame `json:"games"`
}

// possible improvement:
// user <gameid>_mini_boxscore endpoint to have game update self
type NBAGame struct {
	Id        string    `json:"gameId"`
	StartTime time.Time `json:"startTimeUTC"`
	Visitor   NBATeam   `json:"vTeam"`
	Home      NBATeam   `json:"hTeam"`
	Period    Period    `json:"period"`
	Active    bool      `json:"isGameActivated"`
}

type NBABoxScore struct {
	Game *NBAGame `json:"basicGameData"`
}

type NBATeam struct {
	TriCode string `json:"triCode"`
}

type Period struct {
	Current int
}

func (g *NBAGame) GameCode() string {
	return fmt.Sprintf("%s%s", g.Visitor.TriCode, g.Home.TriCode)
}

func (g *NBAGame) IsActive() bool {
	return g.Active
}

func (g *NBAGame) Refresh() {
	resp, err := http.Get(g.gameURL())
	if err != nil {
		log.Printf("error refreshing game: %s\n", err) // TODO
		return
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	var updatedGame NBABoxScore
	for dec.More() {
		err := dec.Decode(&updatedGame)
		if err != nil {
			log.Printf("error decoding game: %s\n", err) // TODO
			return
		}
	}

	if updatedGame.Game == nil {
		return
	}

	g.Active = updatedGame.Game.Active
	g.Period.Current = updatedGame.Game.Period.Current

}

func (g *NBAGame) Plays() []play.Play {
	resp, err := http.Get(g.pbpURL())
	if err != nil {
		log.Printf("error retrieving plays: %s\n", err) // TODO
		return []play.Play{}
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	var plays NBAPlays
	for dec.More() {
		err := dec.Decode(&plays)
		if err != nil {
			log.Printf("error decoding plays: %s\n", err) // TODO
			return []play.Play{}
		}
	}
	g.hydrateNBAPlays(plays)

	return convertNBAPlaysToIPlays(plays) // FIXME -- necessary?
}

func (g *NBAGame) hydrateNBAPlays(plays NBAPlays) {
	for _, play := range plays.Plays {
		play.Visitor = g.Visitor
		play.Home = g.Home
		play.Period = g.Period.Current
	}
}

func (g *NBAGame) gameURL() string {
	return fmt.Sprintf("%s%s", BaseURL, g.gamePath())
}

func (g *NBAGame) gamePath() string {
	const gamePath = "/prod/v1/%s/%s_mini_boxscore.json"
	return fmt.Sprintf(gamePath, g.gameDate(), g.Id)
}

func (g *NBAGame) pbpURL() string {
	return fmt.Sprintf("%s%s", BaseURL, g.pbpPath())
}

func (g *NBAGame) pbpPath() string {
	const pbpPath = "/prod/v1/%s/%s_pbp_%d.json"
	return fmt.Sprintf(pbpPath, g.gameDate(), g.Id, g.Period.Current)
}

// GameDate returns the start date of game (YYYYMMDD format) in US/Eastern tz
func (g *NBAGame) gameDate() string {
	easternTime, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic("getting game date")
	}
	return g.StartTime.In(easternTime).Format("20060102")
}

func (g *NBAGame) String() string {
	return g.GameCode()
}

func convertNBAPlaysToIPlays(plays NBAPlays) []play.Play {
	var iPlays = make([]play.Play, len(plays.Plays))
	for i, play := range plays.Plays {
		iPlays[i] = play
	}
	return iPlays
}

type NBAPlays struct {
	Plays []*NBAPlay `json:"plays"`
}

type NBAPlay struct {
	Clock            string           `json:"clock"`
	Description      string           `json:"description"`
	PersonId         string           `json:"personId"`
	TeamId           string           `json:"teamId"`
	VistingTeamScore string           `json:"vTeamScore"`
	HomeTeamScore    string           `json:"hTeamScore"`
	Formatted        FormattedNBAPlay `json:"formatted"`
	Visitor          NBATeam
	Home             NBATeam
	Period           int
}

type FormattedNBAPlay struct {
	Description string `json:"description"`
}

func (p NBAPlay) String() string {
	return fmt.Sprintf("%s\n[%s %s - %s %s]\n[%s Q%d]", p.Formatted.Description, p.Visitor.TriCode, p.VistingTeamScore, p.HomeTeamScore, p.Home.TriCode, p.Clock, p.Period)
}
