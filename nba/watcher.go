package nba

import (
	_ "fmt"

	"github.com/andrewmelis/nba-tweeter/game"
	"github.com/andrewmelis/nba-tweeter/schedule"
)

type NBAWatcher struct {
	p *Processor
}

func NewNBAWatcher(p *Processor) NBAWatcher {
	return NBAWatcher{p}
}

func (w *NBAWatcher) Follow(schedule.Schedule) {
	// TODO
}

func (w *NBAWatcher) IsWatching(g game.Game) bool {
	return w.p.IsProcessing(g) // FIXME
}

type Processor struct {
}

func (p *Processor) IsProcessing(g game.Game) bool {
	return false
}
