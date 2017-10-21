package nba

import (
	"testing"
	"time"

	"github.com/andrewmelis/nba-tweeter/game"
	"github.com/andrewmelis/nba-tweeter/watcher"
)

func TestFollowStartsWatcherForEachActiveGame(t *testing.T) {
	setupNow()
	advanceTimeCh := setupTicker()

	games := fakeNBAGames("OKCBOS", "GSWCLE", "SASHOU")
	games.Games[len(games.Games)-1].Active = false // set last one to be inactive

	s := newFakeSchedule(games)
	f := NewNBAFollower()

	hookCh := make(chan struct{})
	followHook = func() { <-hookCh }

	f.Follow(s)

	advanceTimeCh <- Now().Add(10 * time.Second)
	hookCh <- struct{}{}

	for i, game := range games.Games {
		w, ok := f.watchedGames.Load(game.GameCode())

		if i < len(games.Games)-1 {
			if !ok {
				t.Errorf("Expected watcher for game %#v!", game)
			}

			if _, ok = w.(watcher.Watcher); !ok {
				t.Errorf("Expected Watcher type as value in watchedGames map!")
			}
		} else { // last game is not yet active
			if ok {
				t.Errorf("Expected no watcher for game %#v!", game)
			}
		}
	}

	// advance time, ensure new set of active games included
	secondSetGames := fakeNBAGames("OKCBOS", "GSWCLE", "SASHOU") // all active
	s.games = secondSetGames                                     // intentionally replace references

	advanceTimeCh <- Now().Add(10 * time.Second)
	hookCh <- struct{}{}

	for _, game := range secondSetGames.Games {
		w, ok := f.watchedGames.Load(game.GameCode())
		if !ok {
			t.Errorf("Expected watcher for game %#v!", game)
		}

		if _, ok = w.(watcher.Watcher); !ok {
			t.Errorf("Expected Watcher type as value in watchedGames map!")
		}
	}

	// advance time again, ensure finished games are removed
	thirdSetGames := fakeNBAGames("OKCBOS", "GSWCLE", "SASHOU")
	thirdSetGames.Games[0].Active = false // set last one to be inactive
	s.games = thirdSetGames               // intentionally replace references

	advanceTimeCh <- Now().Add(10 * time.Second)
	hookCh <- struct{}{}

	for i, game := range thirdSetGames.Games {
		w, ok := f.watchedGames.Load(game.GameCode())

		if i == 0 { // first game is now inactive
			if ok {
				t.Errorf("Expected no watcher for game %#v!", game)
			}
		} else {
			if !ok {
				t.Errorf("Expected watcher for game %#v!", game)
			}

			if _, ok = w.(watcher.Watcher); !ok {
				t.Errorf("Expected Watcher type as value in watchedGames map!")
			}
		}
	}
}

type fakeSchedule struct {
	games NBAGames
}

func newFakeSchedule(games NBAGames) *fakeSchedule {
	return &fakeSchedule{games}
}

func (f *fakeSchedule) Games() []game.Game {
	return convertNBAGamesToIGames(f.games)
}
