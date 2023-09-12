package postgresql

import (
	"database/sql"
	"fmt"
	"go-rest-api/entities"

	_ "github.com/lib/pq"
)

var connStr = "user=postgres dbname=gamedatabase sslmode=disable"

func InsertNewPlayer(tableName string, playerId string, playername string, score string) error {
	query := fmt.Sprintf("INSERT INTO %s (player_id, player_name, score) VALUES ('%s', '%s', '%s')", tableName, playerId, playername, score)

	err := interactWithDb(query)
	if err != nil {
		return err
	}

	return nil
}

func DeletePlayer(tableName string, playerId string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE player_id = '%s'", tableName, playerId)

	err := interactWithDb(query)
	if err != nil {
		return err
	}

	return nil
}

func UpdatePlayerScore(tableName string, playerId string, newScore string) error {
	query := fmt.Sprintf("UPDATE %s SET score = '%s' WHERE player_id = '%s'", tableName, newScore, playerId)

	err := interactWithDb(query)
	if err != nil {
		return err
	}

	return nil
}

func UpdatePlayerName(tableName string, playerId string, newName string) error {
	query := fmt.Sprintf("UPDATE %s SET player_name = '%s' WHERE player_id = '%s'", tableName, newName, playerId)

	err := interactWithDb(query)
	if err != nil {
		return err
	}

	return nil
}

func GetPlayerInfo(tableName string, playerId string) (*entities.PlayerInfo, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT player_name, score FROM %s WHERE player_id = '%s'", tableName, playerId)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var playerName, score string
		err := rows.Scan(&playerName, &score)
		if err != nil {
			return nil, err
		}

		return &entities.PlayerInfo{
			PlayerName: playerName,
			Score:      score,
		}, nil
	}

	return nil, fmt.Errorf("User not found")
}

func GetLeaderboard(tableName string) ([]entities.PlayerInfo, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT * FROM %s ORDER BY CAST(score AS INT) DESC LIMIT 10", tableName)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var response []entities.PlayerInfo
	for rows.Next() {
		var id, playerName, score string

		err := rows.Scan(&id, &playerName, &score)
		if err != nil {
			return nil, err
		}
		response = append(response, entities.PlayerInfo{
			PlayerName: playerName,
			Score:      score,
		})
	}

	if len(response) > 0 {
		return response, nil
	} else {
		return nil, fmt.Errorf("No records found")
	}

}

func interactWithDb(query string) error {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
