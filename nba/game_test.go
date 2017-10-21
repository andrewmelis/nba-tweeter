package nba

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGameCode(t *testing.T) {
	g := newNBAGame("WASCLE")

	expected := "WASCLE"
	actual := g.GameCode()

	if expected != actual {
		t.Errorf("expected: %s, actual: %s\n", expected, actual)
	}
}

func newNBAGame(code string) *NBAGame {
	visitorCode := code[:3]
	visitor := NBATeam{visitorCode}

	homeCode := code[3:]
	home := NBATeam{homeCode}

	return &NBAGame{
		Id:        "0041600404",
		StartTime: makeTime("20170609 7:30pm", "US/Eastern"),
		Visitor:   visitor,
		Home:      home,
		Active:    true,
		Period: Period{
			Current: 1,
		},
	}
}

func TestGameRefresh(t *testing.T) {
	ts := httptest.NewServer(newPbpFixtureHandlerFunc())
	defer ts.Close()

	BaseURL = ts.URL

	g := newNBAGame("WASCLE")
	g.Refresh()

	if g.Period.Current != 4 {
		t.Errorf("Expected game period to be updated to %d, got %d", 4, g.Period.Current)
	}
	if g.Active {
		t.Errorf("Expected game to be updated to inactive, got %t", g.Active)
	}
}

func TestPlaysRetrieval(t *testing.T) {
	ts := httptest.NewServer(newPbpFixtureHandlerFunc())
	defer ts.Close()

	BaseURL = ts.URL

	g := newNBAGame("GSWCLE")

	expectedPlays := []string{
		`CLE - Williams Turnover : Lost Ball (1 TO) Steal:Green (2 ST)
[GSW 96 - 115 CLE]
[11:50 Q1]`,
		`GSW - McCaw 3pt Shot: Made (3 PTS) Assist: K Thompson (2 AST)
[GSW 99 - 115 CLE]
[11:21 Q1]`,
	}
	actualPlays := g.Plays()

	if len(expectedPlays) != len(actualPlays) {
		t.Errorf("Wanted: %d, Got: %d", len(expectedPlays), len(actualPlays))
	}

	for i, play := range actualPlays {
		expectedPlay := expectedPlays[i]
		actualPlay := play.String()
		if expectedPlay != actualPlay {
			t.Errorf("Wanted: %s, Got: %s", expectedPlay, actualPlay)
		}
	}
}

func newPbpFixtureHandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var filename string

		switch r.URL.Path {
		case "/prod/v1/20170609/0041600404_pbp_1.json":
			filename = "fixtures/pbp.json"
		case "/prod/v1/20170609/0041600404_mini_boxscore.json":
			filename = "fixtures/mini_boxscore.json"
		}
		contents, err := ioutil.ReadFile(filename)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, `{"error":"server error occurred"}`)
			return
		}
		fmt.Fprintf(w, "%s", contents)
	}
}
