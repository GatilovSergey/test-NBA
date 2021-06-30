package nba


type GameInfoPerMin struct{
	Minutes int
	GameInfo
}

type GameInfo struct {
	First Team
	Second Team
}

type Team struct {
	Attack int
	Score int
}

func NewGameInfo() GameInfo {
	return GameInfo{Team{0,0},Team{0,0}}
}
