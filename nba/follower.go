package nba

import (
	"sync"
	"time"

	"github.com/andrewmelis/nba-tweeter/game"
	"github.com/andrewmelis/nba-tweeter/processor"
	"github.com/andrewmelis/nba-tweeter/schedule"
)

type NBAFollower struct {
	watchedGames sync.Map
	followHook   func()
}

func NewNBAFollower() *NBAFollower {
	m := sync.Map{}
	return &NBAFollower{
		watchedGames: m,
		followHook:   func() {},
	}
}

func (f *NBAFollower) Follow(s schedule.Schedule) {
	go func() {
		t := NewTicker(10 * time.Second)
		for range t.C {
			for _, game := range s.Games() {
				switch {
				case game.IsActive() && !f.isBeingWatched(game):
					f.watchGame(game)
				case !game.IsActive() && f.isBeingWatched(game):
					f.removeGame(game) // probably should use callback from watcher
				}
			}
			f.followHook()
		}
	}() // should i be passing args in to this anon func?
}

func (f *NBAFollower) watchGame(g game.Game) {
	p := processor.NewDebugProcessor() // FIXME
	w := NewNBAWatcher(p, func(string) {})

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
