package helper

import (
	"os"
	"testing"

	"fmt"

	. "github.com/franela/goblin"
)

var jsonFilePath = "./quake_data_temp.json"

func TestHelperCheckError(t *testing.T) {
	g := Goblin(t)

	g.Describe("CHECKERROR", func() {
		g.It("Should return TRUE", func() {
			ok := CheckError(fmt.Errorf("Some ERROR!"))

			g.Assert(ok).IsTrue()
		})

		g.It("Should return FALSE", func() {
			ok := CheckError(nil)

			g.Assert(ok).IsFalse()
		})
	})
}

func TestHelperGetGames(t *testing.T) {
	g := Goblin(t)

	g.Describe("GETGAMES", func() {
		g.Before(func() {
			createQuakeTempData()
		})

		g.After(func() {
			deleteQuakeTempData()
		})

		g.It("Should return the games and no error", func() {
			games, err := GetGames(jsonFilePath)

			expectedGame := &Game{
				TotalKills: 0,
				Players: map[int]string{
					1: "Isgalamido",
				},
				Kills: map[string]int{
					"Isgalamido": 0,
				},
			}

			g.Assert(games).Equal(map[string]*Game{"game_1": expectedGame})
			g.Assert(err).Equal(nil)
		})

		g.It("Should return no games and a error", func() {
			games, err := GetGames("")

			g.Assert(games).Equal(map[string]*Game{})
			g.Assert(CheckError(err)).IsTrue()
		})
	})
}

func createQuakeTempData() {
	jsonStr := `{"game_1": {
		"total_kills": 0,
		"players": {
			"1": "Isgalamido"
		},
		"kills": {
			"Isgalamido": 0
		}
	}
}`

	f, err := os.Create(jsonFilePath)
	if err != nil {
		panic(err)
	}

	_, err = f.WriteString(jsonStr)
	if err != nil {
		panic(err)
	}
}

func deleteQuakeTempData() {
	err := os.Remove(jsonFilePath)
	if err != nil {
		panic(err)
	}
}
