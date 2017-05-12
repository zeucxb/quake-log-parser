package helper

import (
	"encoding/json"
	"io/ioutil"
)

// Game represents the parsed data
type Game struct {
	TotalKills int            `json:"total_kills"`
	Players    map[int]string `json:"players"`
	Kills      map[string]int `json:"kills"`
}

// GetGames open the json with the received path, populate games and return
func GetGames(jsonFilePath string) (games map[string]*Game, err error) {
	jsonFile, err := ioutil.ReadFile(jsonFilePath)

	games = make(map[string]*Game)

	err = json.Unmarshal(jsonFile, &games)

	return
}

// CheckError receives a error and check it
func CheckError(err error) (ok bool) {
	if err != nil {
		return true
	}

	return
}
