package nba

import (
	"encoding/json"
	"fmt"
	_ "io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andrewmelis/nba-tweeter/game"
)

func TestScheduledNBAGames(t *testing.T) {
	var schedules = [][]game.Game{
		NewNBAGames("GSWCLE"),
		NewNBAGames("ATLWAS"),
		NewNBAGames("GSWCLE", "ATLWAS"),
	}

	for _, s := range schedules {
		testSchedule(t, s)
	}
}

func NewNBAGames(gameCodes ...string) []game.Game {
	var games []game.Game
	for _, code := range gameCodes {
		games = append(games, NewNBAGame(code))
	}
	return games
}

func testSchedule(t *testing.T, expectedGames []game.Game) {
	ts := httptest.NewServer(newNBAGameHandlerFunc(expectedGames))
	defer ts.Close()

	r := NewFakeScheduleURL(ts.URL)
	s := NewNBASchedule(r) // inject fakes into constructor

	actualGames := s.ScheduledGames()

	for i := range expectedGames {
		expected := expectedGames[i]
		actual := actualGames[i]
		if expected != actual {
			t.Errorf("expected: %+v; actual: %+v\n", expectedGames, actualGames)
		}
	}
}

func newNBAGameHandlerFunc(g []game.Game) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)
		err := enc.Encode(&g)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, `{"error":"server error occurred"}`)
			return
		}
	}
}

type FakeScheduleURL struct {
	url string
}

func NewFakeScheduleURL(url string) FakeScheduleURL {
	return FakeScheduleURL{url: url}
}

func (r FakeScheduleURL) URL() string {
	return r.url
}
