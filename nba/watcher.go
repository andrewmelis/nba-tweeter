package nba

import (
	_ "log"
	"time"

	"github.com/andrewmelis/nba-tweeter/game"
	"github.com/andrewmelis/nba-tweeter/play"
	"github.com/andrewmelis/nba-tweeter/processor"
)

type NBAWatcher struct {
	p         processor.Processor
	seenPlays []play.Play
	cb        func(string)
}

func NewNBAWatcher(p processor.Processor, cb func(string)) *NBAWatcher {
	return &NBAWatcher{
		p:         p,
		seenPlays: make([]play.Play, 0),
		cb:        cb,
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

			inputs := g.Plays()
			inputs = unique(inputs)

			for i, play := range inputs {
				if i >= len(w.seenPlays) {
					w.seenPlays = append(w.seenPlays, play)
					w.p.Process(code, play)
				}
			}

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
