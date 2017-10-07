package nba

import (
	_ "fmt"

	"github.com/andrewmelis/nba-tweeter/clock"
	"github.com/andrewmelis/nba-tweeter/processor"
	"github.com/andrewmelis/nba-tweeter/schedule"
)

type NBAWatcher struct {
	c clock.Clock
	p processor.Processor
}

func NewNBAWatcher(c clock.Clock, p processor.Processor) NBAWatcher {
	return NBAWatcher{c, p}
}

func (w *NBAWatcher) Follow(schedule.Schedule) {

	// TODO
}
