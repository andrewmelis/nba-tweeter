package nba

import (
	_ "fmt"

	"github.com/andrewmelis/nba-tweeter/clock"
	"github.com/andrewmelis/nba-tweeter/game"
	"github.com/andrewmelis/nba-tweeter/processor"
)

type NBAWatcher struct {
	c clock.Clock
	p processor.Processor
}

func NewNBAWatcher(c clock.Clock, p processor.Processor) NBAWatcher {
	return NBAWatcher{c, p}
}

func (w *NBAWatcher) Follow(g game.Game) {
	for _, play := range g.Plays() {
		w.p.Process(g.GameCode(), play)
	}
}
