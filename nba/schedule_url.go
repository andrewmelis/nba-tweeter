package nba

import (
	"fmt"
)

type ScheduleURL interface {
	URL() string
}

type NBAScheduleURL struct {
	baseURL string
}

func NewNBAScheduleURL(baseURL string) ScheduleURL {
	return NBAScheduleURL{baseURL}
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
	return Now().Format("20060102")
}
