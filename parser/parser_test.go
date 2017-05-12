package parser

import (
	"os"
	"testing"

	. "github.com/franela/goblin"
)

var logFilePath = "./games_temp.log"

func createTempLog() {
	logStr := `0:00 ------------------------------------------------------------
  0:00 InitGame: \sv_floodProtect\1\sv_maxPing\0\sv_minPing\0\sv_maxRate\10000\sv_minRate\0\sv_hostname\Code Miner Server\g_gametype\0\sv_privateClients\2\sv_maxclients\16\sv_allowDownload\0\dmflags\0\fraglimit\20\timelimit\15\g_maxGameClients\0\capturelimit\8\version\ioq3 1.36 linux-x86_64 Apr 12 2009\protocol\68\mapname\q3dm17\gamename\baseq3\g_needpass\0
 15:00 Exit: Timelimit hit.
 20:34 ClientConnect: 2
 20:34 ClientUserinfoChanged: 2 n\Isgalamido\t\0\model\xian/default\hmodel\xian/default\g_redteam\\g_blueteam\\c1\4\c2\5\hc\100\w\0\l\0\tt\0\tl\0
 20:37 ClientUserinfoChanged: 2 n\Isgalamido\t\0\model\uriel/zael\hmodel\uriel/zael\g_redteam\\g_blueteam\\c1\5\c2\5\hc\100\w\0\l\0\tt\0\tl\0
 20:37 ClientBegin: 2
 20:37 ShutdownGame:
 20:37 ------------------------------------------------------------`

	f, err := os.Create(logFilePath)
	if err != nil {
		panic(err)
	}

	_, err = f.WriteString(logStr)
	if err != nil {
		panic(err)
	}
}

func deleteTempLogAndGeneratedFile() {
	err := os.Remove(logFilePath)
	if err != nil {
		panic(err)
	}

	err = os.Remove("quake_data.json")
	if err != nil {
		panic(err)
	}
}

func TestParserParse(t *testing.T) {
	g := Goblin(t)

	g.Describe("Parse", func() {
		g.BeforeEach(func() {
			createTempLog()
		})

		g.AfterEach(func() {
			deleteTempLogAndGeneratedFile()
		})

		g.It("Should receive a valid file and return no errors", func() {
			err := Parse(logFilePath)

			g.Assert(err).Equal(nil)
		})

		g.It("Should parse the file correctly", func() {
			Parse(logFilePath)

			game, ok := games["game_1"]

			g.Assert(ok).IsTrue()
			g.Assert(game.TotalKills).Equal(0)
			g.Assert(game.Players).Equal(map[int]string{1: "Isgalamido"})
			g.Assert(game.Kills).Equal(map[string]int{"Isgalamido": 0})
		})

		g.It("Should create a file", func() {
			Parse(logFilePath)

			_, err := os.Stat(logFilePath)

			g.Assert(err).Equal(nil)
		})
	})
}

func TestParserInitRegEx(t *testing.T) {
	g := Goblin(t)

	g.Describe("initRegEx", func() {
		g.It("Should init regex and return no errors", func() {
			err := initRegEx()

			g.Assert(err).Equal(nil)
		})
	})
}
