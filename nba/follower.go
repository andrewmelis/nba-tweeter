package nba

import (
	_ "log"
	"sync"
	"time"

	"github.com/andrewmelis/nba-tweeter/game"
	"github.com/andrewmelis/nba-tweeter/processor"
	"github.com/andrewmelis/nba-tweeter/schedule"
)

type NBAFollower struct {
	watchedGames sync.Map
	pf           processor.ProcessorFactory
}

func NewNBAFollower(pf processor.ProcessorFactory) *NBAFollower {
	m := sync.Map{}
	return &NBAFollower{
		watchedGames: m,
		pf:           pf,
	}
}

var followHook func() = func() {}

func (f *NBAFollower) Follow(s schedule.Schedule) {
	go func() {
		t := NewTicker(10 * time.Second)
		for range t.C {
			// log.Printf("Follow: retrieving schedule")
			for _, game := range s.Games() {
				// log.Printf("Follow: switching game %s", game.GameCode())
				switch {
				case game.IsActive() && !f.isBeingWatched(game):
					// log.Printf("Follow: start watching game %s", game.GameCode())
					f.watchGame(game)
				case !game.IsActive() && f.isBeingWatched(game):
					// log.Printf("Follow: removing game %s", game.GameCode())
					f.removeGame(game) // probably should use callback from watcher
				}
			}
			// log.Printf("Follow: firing followHook")
			followHook()
			// log.Printf("Follow: after firing followHook")
		}
	}() // should i be passing args in to this anon func?
}

func (f *NBAFollower) watchGame(g game.Game) {
	p := f.pf(g)
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
