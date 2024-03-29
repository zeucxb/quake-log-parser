package parser

import (
	"fmt"
	"quake-log-parser/helper"
	"strconv"
	"strings"
)

func checkAndParseKill(text string) (check bool, err error) {
	if match := killRegEx.FindStringSubmatch(text); len(match) == 2 {
		killInfo := strings.Split(match[1], " ")

		killerKey, err := strconv.Atoi(killInfo[0])
		if err != nil {
			return false, err
		}

		killedKey, err := strconv.Atoi(killInfo[1])
		if err != nil {
			return false, err
		}

		if killInfo[0] == "1022" {
			games[key].Kills[games[key].Players[killedKey-1]]--
		} else {
			games[key].Kills[games[key].Players[killerKey-1]]++
		}

		games[key].TotalKills++

		return true, err
	}

	return
}

func checkAndParseUser(text string) (check bool, err error) {
	if match := userRegEx.FindStringSubmatch(text); len(match) == 3 {
		playerKey, err := strconv.Atoi(match[1])
		if err != nil {
			return false, err
		}

		if _, ok := games[key].Players[playerKey-1]; !ok {
			newPlayer := match[2]

			games[key].Players[playerKey-1] = newPlayer
			games[key].Kills[newPlayer] = 0
		} else {
			currentPlayer := games[key].Players[playerKey-1]
			currentKills := games[key].Kills[currentPlayer]

			newPlayer := match[2]

			delete(games[key].Kills, currentPlayer)

			games[key].Players[playerKey-1] = newPlayer
			games[key].Kills[newPlayer] = currentKills
		}

		return true, nil
	}

	return
}

func checkAndParseGameInit(text string) (check bool, err error) {
	if match := gameStartRegEx.MatchString(text); match {
		key = fmt.Sprintf("%s%v", keyPref, count)

		if _, ok := games[key]; !ok {
			games[key] = &helper.Game{
				Players: make(map[int]string),
				Kills:   make(map[string]int),
			}
		}

		count++

		return true, nil
	}

	return
}

func checkAndParse(text string) (err error) {
	if check, err := checkAndParseKill(text); check {
		return err
	}

	if check, err := checkAndParseUser(text); check {
		return err
	}

	if check, err := checkAndParseGameInit(text); check {
		return err
	}

	return
}
