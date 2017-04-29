package nba

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	actualGames := s.Games()

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

func TestExampleJSON(t *testing.T) {
	ts := httptest.NewServer(newFixtureHandlerFunc("fixtures/example.json"))
	defer ts.Close()

	r := newFakeScheduleURL(ts.URL)
	s := NewNBASchedule(r) // inject fakes into constructor

	expectedCodes := []string{"WASBOS", "UTALAC"}

	actualGames := s.Games()

	for i := range expectedCodes {
		expected := expectedCodes[i]
		actual := actualGames[i].GameCode()
		if expected != actual {
			t.Errorf("expected: %+v; actual: %+v\n", expected, actual)
		}
	}
}

func newFixtureHandlerFunc(filename string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contents, err := ioutil.ReadFile(filename)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, `{"error":"server error occurred"}`)
			return
		}
		fmt.Fprintf(w, "%s", contents)
	}
}
