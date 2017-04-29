package nba

type ScheduleURL interface {
	URL() string
}

type NBAScheduleURL struct{}

func NewDefaultNBAScheduleURL() ScheduleURL {
	return NBAScheduleURL{}
}

func (u NBAScheduleURL) URL() string {
	return "https://data.nba.net/data/10s/prod/v1/20170430/scoreboard.json"
}
