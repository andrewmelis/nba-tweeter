package nba

type NBAGame struct {
	Code string `json:"code"`
}

func NewNBAGame(code string) NBAGame {
	return NBAGame{Code: code}
}

func (g NBAGame) GameCode() string {
	return g.Code
}
