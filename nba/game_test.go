package nba

import (
	"testing"
)

func TestGameCode(t *testing.T) {
	g := fakeNBAGame("WASCLE")

	expected := "WASCLE"
	actual := g.GameCode()

	if expected != actual {
		t.Errorf("expected: %s, actual: %s\n", expected, actual)
	}
}

func fakeNBAGame(code string) NBAGame {
	visitorCode := code[:3]
	visitor := NBATeam{visitorCode}

	homeCode := code[3:]
	home := NBATeam{homeCode}

	return NBAGame{visitor, home}
}
