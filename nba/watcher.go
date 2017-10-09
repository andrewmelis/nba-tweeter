package nba

import (
	_ "fmt"

	"github.com/andrewmelis/nba-tweeter/clock"
	"github.com/andrewmelis/nba-tweeter/game"
	"github.com/andrewmelis/nba-tweeter/play"
	"github.com/andrewmelis/nba-tweeter/processor"
)

type NBAWatcher struct {
	c         clock.Clock
	p         processor.Processor
	seenPlays map[string][]play.Play
	cb        func(string)
}

func NewNBAWatcher(c clock.Clock, p processor.Processor, cb func(string)) NBAWatcher {
	return NBAWatcher{
		c:         c,
		p:         p,
		seenPlays: make(map[string][]play.Play),
		cb:        cb,
	}
}

func (w *NBAWatcher) Follow(g game.Game) {
	go func() {
		for range w.c.Ticker() {
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

			if !gameStillActive {
				w.cb(code)
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
