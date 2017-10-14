package nba

import (
	"fmt"

	"github.com/andrewmelis/nba-tweeter/clock"
)

type ScheduleURL interface {
	URL() string
}

type NBAScheduleURL struct {
	baseURL string
	c       clock.Clock
}

func NewNBAScheduleURL(baseURL string, c clock.Clock) ScheduleURL {
	return NBAScheduleURL{baseURL, c}
}

func (u NBAScheduleURL) URL() string {
	// TODO retrieve from "today url":
	// https://data.nba.net/10s/prod/v1/today.json -> currentScoreboard

	// vvvvvvvvvvvvvvvvvvvvvvvv
	// https://data.nba.net/
	// ^^^^^^^^^^^^^^^^^^^^^^^^

	return fmt.Sprintf("%s%s", u.baseURL, u.scoreboardPath())
}

func (u NBAScheduleURL) scoreboardPath() string {
	const scoreboardPath = "/data/10s/prod/v1/%s/scoreboard.json"
	return fmt.Sprintf(scoreboardPath, u.date())
}

func (u NBAScheduleURL) date() string {
	return u.c.Now().Format("20060102")
}
