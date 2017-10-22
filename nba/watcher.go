package nba

import (
	"log"
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
			log.Printf("Refreshing game: %s\n", g)
			g.Refresh()

			code := g.GameCode()
			gameStillActive := g.IsActive()

			w.processNewPlays(code, g.Plays())

			WatchHook()
			if !gameStillActive {
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
