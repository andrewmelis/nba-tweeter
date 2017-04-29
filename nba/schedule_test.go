package nba

import (
	"encoding/json"
	"fmt"
	_ "io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestScheduledNBAGames(t *testing.T) {
	var schedules = []NBAGames{
		fakeNBAGames("GSWCLE"),
		fakeNBAGames("GSWCLE", "ATLWAS"),
		fakeNBAGames("GSWCLE", "ATLWAS", "CHIBOS"),
	}

	for _, s := range schedules {
		testSchedule(t, s)
	}
}

func testSchedule(t *testing.T, expectedGames NBAGames) {
	ts := httptest.NewServer(newNBAGameHandlerFunc(expectedGames))
	defer ts.Close()

	r := newFakeScheduleURL(ts.URL)
	s := NewNBASchedule(r) // inject fakes into constructor

	actualGames := s.ScheduledGames()

	for i := range expectedGames.Games {
		expected := expectedGames.Games[i]
		actual := actualGames[i]
		if expected != actual {
			t.Errorf("expected: %+v; actual: %+v\n", expectedGames, actualGames)
		}
	}
}

func fakeNBAGames(gameCodes ...string) NBAGames {
	var games NBAGames
	for _, code := range gameCodes {
		games.Games = append(games.Games, fakeNBAGame(code))
	}
	return games
}

func newNBAGameHandlerFunc(g NBAGames) http.HandlerFunc {
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

type fakeScheduleURL struct {
	url string
}

func newFakeScheduleURL(url string) fakeScheduleURL {
	return fakeScheduleURL{url: url}
}

func (r fakeScheduleURL) URL() string {
	return r.url
}
