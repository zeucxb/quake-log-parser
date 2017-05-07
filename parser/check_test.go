package parser

import (
	"testing"

	. "github.com/franela/goblin"
)

func Test(t *testing.T) {
	g := Goblin(t)

	initRegEx()

	gameInitText := `0:00 InitGame: \sv_floodProtect\1\sv_maxPing\0\sv_minPing\0\sv_maxRate\10000\sv_minRate\0\sv_hostname\Code Miner Server\g_gametype\0\sv_privateClients\2\sv_maxclients\16\sv_allowDownload\0\dmflags\0\fraglimit\20\timelimit\15\g_maxGameClients\0\capturelimit\8\version\ioq3 1.36 linux-x86_64 Apr 12 2009\protocol\68\mapname\q3dm17\gamename\baseq3\g_needpass\0`
	userText := `ClientUserinfoChanged: 2 n\Isgalamido\t\0\model\xian/default\hmodel\xian/default\g_redteam\\g_blueteam\\c1\4\c2\5\hc\100\w\0\l\0\tt\0\tl\0`
	killText := `0:17 Kill: 1022 3 22: <world> killed Dono da Bola by MOD_TRIGGER_HURT`

	g.Describe("GAME INIT parse", func() {
		g.AfterEach(func() {
			count = 1
			key = ""

			games = make(map[string]*game)
		})

		g.It("Should return as checked and has no errors", func() {
			check, err := checkAndParseGameInit(gameInitText)

			g.Assert(check).IsTrue()
			g.Assert(err).Equal(nil)
		})

		g.It("Should return as not checked and has no errors", func() {
			check, err := checkAndParseGameInit(userText)

			g.Assert(check).IsFalse()
			g.Assert(err).Equal(nil)
		})

		g.It("Should update the key and the count", func() {
			checkAndParseGameInit(gameInitText)

			g.Assert(key).Equal("game_1")
			g.Assert(count).Equal(2)
		})

		g.It("Should create a game correctly", func() {
			checkAndParseGameInit(gameInitText)

			game, ok := games["game_1"]

			g.Assert(ok).IsTrue()
			g.Assert(game.TotalKills).Equal(0)
			g.Assert(game.Players).Equal(make(map[int]string))
			g.Assert(game.Kills).Equal(make(map[string]int))
		})
	})

	g.Describe("USER parse", func() {
		g.Before(func() {
			checkAndParseGameInit(gameInitText)
		})

		g.After(func() {
			games = make(map[string]*game)
		})

		g.It("Should return as checked and has no errors", func() {
			check, err := checkAndParseUser(userText)

			g.Assert(check).IsTrue()
			g.Assert(err).Equal(nil)
		})
	})

	g.Describe("KILL parse", func() {
		g.Before(func() {
			checkAndParseGameInit(gameInitText)
		})

		g.After(func() {
			games = make(map[string]*game)
		})

		g.It("Should return as checked and has no errors", func() {
			check, err := checkAndParseKill(killText)

			g.Assert(check).IsTrue()
			g.Assert(err).Equal(nil)
		})
	})
}
