package nba

import (
	_ "log"
	"time"

	"github.com/andrewmelis/nba-tweeter/game"
	"github.com/andrewmelis/nba-tweeter/play"
	"github.com/andrewmelis/nba-tweeter/processor"
)

type NBAWatcher struct {
	p                processor.Processor
	seenPlays        map[string]bool
	orderedSeenPlays []play.Play
	cb               func(string)
}

func NewNBAWatcher(p processor.Processor, cb func(string)) *NBAWatcher {
	return &NBAWatcher{
		p:                p,
		seenPlays:        make(map[string]bool),
		orderedSeenPlays: make([]play.Play, 0),
		cb:               cb,
	}
}

var WatchHook func() = func() {}

func (w *NBAWatcher) Watch(g game.Game) {
	go func() {
		t := NewTicker(10 * time.Second)
		for range t.C {
			// log.Printf("Watch: refreshing game %#v", g)
			g.Refresh()
			// log.Printf("Watch: refreshed game %#v", g)

			code := g.GameCode()
			gameStillActive := g.IsActive()

			w.processNewPlays(code, g.Plays())

			// log.Printf("Watch: before watchhook")
			WatchHook()
			// log.Printf("Watch: refreshed game %s", g.GameCode())
			if !gameStillActive {
				// log.Printf("Watch: firing game over callback %s", g.GameCode())
				w.cb(code)
				t.Stop()
				return
			}
		}
	}()
}

func (w *NBAWatcher) processNewPlays(gamecode string, in []play.Play) {
	for _, candidate := range in {
		if w.seenPlays[candidate.String()] {
			continue
		}
		w.orderedSeenPlays = append(w.orderedSeenPlays, candidate)
		w.seenPlays[candidate.String()] = true
		w.p.Process(gamecode, candidate)
	}
}
