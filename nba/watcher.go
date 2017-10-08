package nba

import (
	_ "fmt"

	"github.com/andrewmelis/nba-tweeter/clock"
	"github.com/andrewmelis/nba-tweeter/game"
	"github.com/andrewmelis/nba-tweeter/play"
	"github.com/andrewmelis/nba-tweeter/processor"
)

type NBAWatcher struct {
	c         clock.Clock
	p         processor.Processor
	seenPlays []play.Play
}

func NewNBAWatcher(c clock.Clock, p processor.Processor) NBAWatcher {
	return NBAWatcher{
		c:         c,
		p:         p,
		seenPlays: make([]play.Play, 0),
	}
}

func (w *NBAWatcher) Follow(g game.Game) {
	inputs := g.Plays()
	inputs = unique(inputs)

	for i, play := range inputs {
		if i >= len(w.seenPlays) {
			w.seenPlays = append(w.seenPlays, play)
			w.p.Process(g.GameCode(), play)
		}
	}
}

func unique(in []play.Play) []play.Play {
	found := make(map[string]bool)
	out := make([]play.Play, 0)

	for _, candidate := range in {
		if found[candidate.String()] {
			continue
		}
		found[candidate.String()] = true
		out = append(out, candidate)
	}
	return out
}
