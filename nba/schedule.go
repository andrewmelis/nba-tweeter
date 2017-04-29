package nba

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/andrewmelis/nba-tweeter/game"
)

type NBASchedule struct {
	r ScheduleURL
}

func NewDefaultNBASchedule() *NBASchedule {
	return &NBASchedule{r: NewDefaultNBAScheduleURL()}
}

func NewNBASchedule(r ScheduleURL) *NBASchedule {
	return &NBASchedule{r: r}
}

// ScheduledGames retrieves, decodes, and returns a slice of Game types
// from the calling *NBASchedule's ScheduleURL.
func (s *NBASchedule) ScheduledGames() []game.Game {
	resp, err := http.Get(s.r.URL())
	if err != nil {
		log.Printf("error retrieving games: %s\n", err) // TODO
		return []game.Game{}
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	var games []NBAGame
	for dec.More() {
		err := dec.Decode(&games)
		if err != nil {
			log.Printf("error decoding games: %s\n", err) // TODO
			return []game.Game{}
		}
	}

	return convertNBAGamesToIGames(games)
}

func convertNBAGamesToIGames(games []NBAGame) []game.Game {
	var iGames = make([]game.Game, len(games))
	for i, game := range games {
		iGames[i] = game
	}
	return iGames
}
