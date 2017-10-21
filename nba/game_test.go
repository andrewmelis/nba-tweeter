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

	return &NBAGame{Visitor: visitor, Home: home, Active: true}
}

func TestPlaysRetrieval(t *testing.T) {
	ts := httptest.NewServer(newPbpFixtureHandlerFunc())
	defer ts.Close()

	baseURL = ts.URL

	g := newNBAGame("GSWCLE")
	g.Period.Current = 4
	g.StartTime = makeTime("20170609 7:30pm", "US/Eastern")
	g.Id = "0041600404"

	expectedPlays := []string{
		`CLE - Williams Turnover : Lost Ball (1 TO) Steal:Green (2 ST)
[GSW 96 - 115 CLE]
[11:50 Q4]`,
		`GSW - McCaw 3pt Shot: Made (3 PTS) Assist: K Thompson (2 AST)
[GSW 99 - 115 CLE]
[11:21 Q4]`,
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
		case "/data/10s/prod/v1/20170609/0041600404_pbp_4.json":
			filename = "fixtures/pbp.json"
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
