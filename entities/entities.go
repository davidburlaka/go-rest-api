package entities

const TableName = "players"

type Player struct {
	PlayerId   string `json:"player_id"`
	PlayerName string `json:"player_name"`
	Score      string `json:"score"`
}

type PlayerId struct {
	PlayerId string `json:"player_id"`
}

type PlayerInfo struct {
	PlayerName string `json:"player_name"`
	Score      string `json:"score"`
}
