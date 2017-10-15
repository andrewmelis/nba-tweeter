package nba

import (
	"sync"

	"github.com/andrewmelis/nba-tweeter/clock"
	"github.com/andrewmelis/nba-tweeter/game"
	"github.com/andrewmelis/nba-tweeter/processor"
	"github.com/andrewmelis/nba-tweeter/schedule"
)

type NBAFollower struct {
	watchedGames sync.Map
	c            clock.Clock
}

func NewNBAFollower(c clock.Clock) *NBAFollower {
	m := sync.Map{}
	return &NBAFollower{m, c}
}

func (f *NBAFollower) Follow(s schedule.Schedule) {
	go func() {
		for range f.c.Ticker() {
			for _, game := range s.Games() {
				switch {
				case game.IsActive() && !f.isBeingWatched(game):
					f.watchGame(game)
				case !game.IsActive() && f.isBeingWatched(game):
					f.removeGame(game) // probably should use callback from watcher
				}
			}
		}
	}() // should i be passing args in to this anon func?
}

func (f *NBAFollower) watchGame(g game.Game) {
	// inject clock somehow
	now := clock.MakeTime("20170609 7:30pm", "US/Eastern")
	clock := clock.NewFakeClock(now)

	p := processor.NewDebugProcessor() // FIXME
	w := NewNBAWatcher(clock, p, func(string) {})

	// need to somehow ensure both start? i.e., transaction?
	w.Watch(g)
	f.watchedGames.Store(g.GameCode(), w)
	// ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
}

func (f *NBAFollower) removeGame(g game.Game) {
	f.watchedGames.Delete(g.GameCode())
}

func (f *NBAFollower) isBeingWatched(g game.Game) bool {
	_, ok := f.watchedGames.Load(g.GameCode())
	return ok
}
