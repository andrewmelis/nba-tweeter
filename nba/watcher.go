package nba

import (
	"time"

	"github.com/andrewmelis/nba-tweeter/game"
	"github.com/andrewmelis/nba-tweeter/play"
	"github.com/andrewmelis/nba-tweeter/processor"
)

type NBAWatcher struct {
	p         processor.Processor
	seenPlays map[string][]play.Play
	cb        func(string)
	watchHook func()
}

func NewNBAWatcher(p processor.Processor, cb func(string)) *NBAWatcher {
	return &NBAWatcher{
		p:         p,
		seenPlays: make(map[string][]play.Play),
		cb:        cb,
		watchHook: func() {},
	}
}

func (w *NBAWatcher) Watch(g game.Game) {
	go func() {
		t := NewTicker(10 * time.Second)
		for range t.C {
			// update game here! // TODO

			code := g.GameCode()
			gameStillActive := g.IsActive()

			inputs := g.Plays()
			inputs = unique(inputs)

			for i, play := range inputs {
				if i >= len(w.seenPlays[code]) {
					w.seenPlays[code] = append(w.seenPlays[code], play)
					w.p.Process(code, play)
				}
			}

			w.watchHook()
			if !gameStillActive {
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
