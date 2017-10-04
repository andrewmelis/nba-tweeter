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

// Games retrieves, decodes, and returns a slice of Game types
// from the calling *NBASchedule's ScheduleURL.
//
// Potential improvement: have this function simply return stored values
// and have something else populate that store in the background
func (s *NBASchedule) Games() []game.Game {
	resp, err := http.Get(s.r.URL())
	if err != nil {
		log.Printf("error retrieving games: %s\n", err) // TODO
		return []game.Game{}
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	var games NBAGames
	for dec.More() {
		err := dec.Decode(&games)
		if err != nil {
			log.Printf("error decoding games: %s\n", err) // TODO
			return []game.Game{}
		}
	}

	return convertNBAGamesToIGames(games)
}

func convertNBAGamesToIGames(games NBAGames) []game.Game {
	var iGames = make([]game.Game, len(games.Games))
	for i, game := range games.Games {
		iGames[i] = game
	}
	return iGames
}
