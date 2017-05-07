package parser

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type game struct {
	TotalKills int            `json:"total_kills"`
	Players    map[int]string `json:"players"`
	Kills      map[string]int `json:"kills"`
}

// Parse the file and create a json with the correct rules
func Parse(fileStr string) (err error) {
	killRegEx, err := regexp.Compile(`Kill:\s(.*?):`)
	if err != nil {
		return
	}

	userRegEx, err := regexp.Compile(`ClientUserinfoChanged:\s(.)\sn\\(.*)\\t\\`)
	if err != nil {
		return
	}

	gameStartRegEx, err := regexp.Compile(`InitGame:`)
	if err != nil {
		return
	}

	file, err := os.Open(fileStr)
	if err != nil {
		return
	}

	games := make(map[string]*game)
	count := 1
	keyPref := "game_"

	var key string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if match := killRegEx.FindStringSubmatch(scanner.Text()); len(match) == 2 {
			killInfo := strings.Split(match[1], " ")

			killerKey, err := strconv.Atoi(killInfo[0])
			if err != nil {
				panic(err)
			}

			killedKey, err := strconv.Atoi(killInfo[1])
			if err != nil {
				panic(err)
			}

			if killInfo[0] == "1022" {
				games[key].Kills[games[key].Players[killedKey-1]]--
			} else {
				games[key].Kills[games[key].Players[killerKey-1]]++
			}

			games[key].TotalKills++

			continue
		}

		if match := userRegEx.FindStringSubmatch(scanner.Text()); len(match) == 3 {
			playerKey, err := strconv.Atoi(match[1])
			if err != nil {
				panic(err)
			}

			games[key].Players[playerKey-1] = match[2]

			continue
		}

		if match := gameStartRegEx.MatchString(scanner.Text()); match {
			key = fmt.Sprintf("%s%v", keyPref, count)

			if _, ok := games[key]; !ok {
				games[key] = &game{
					Players: make(map[int]string),
					Kills:   make(map[string]int),
				}
			}

			count++

			continue
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	json, err := json.Marshal(games)

	err = ioutil.WriteFile("quake_data.json", json, 0644)

	return
}
