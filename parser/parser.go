package parser

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"regexp"
)

var killRegEx, userRegEx, gameStartRegEx *regexp.Regexp
var games = make(map[string]*game)
var count = 1
var keyPref = "game_"
var key string

type game struct {
	TotalKills int            `json:"total_kills"`
	Players    map[int]string `json:"players"`
	Kills      map[string]int `json:"kills"`
}

func initRegEx() (err error) {
	killRegEx, err = regexp.Compile(`Kill:\s(.*?):`)
	if err != nil {
		return
	}

	userRegEx, err = regexp.Compile(`ClientUserinfoChanged:\s(.)\sn\\(.*)\\t\\`)
	if err != nil {
		return
	}

	gameStartRegEx, err = regexp.Compile(`InitGame:`)
	return
}

func writeFile(games map[string]*game) (err error) {
	json, err := json.Marshal(games)

	err = ioutil.WriteFile("quake_data.json", json, 0644)

	return
}

// Parse the file and create a json with the correct rules
func Parse(fileStr string) (err error) {
	if err := initRegEx(); err != nil {
		return err
	}

	file, err := os.Open(fileStr)
	defer file.Close()
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		checkAndParse(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	err = writeFile(games)

	return
}
