package parser

import (
	"testing"

	. "github.com/franela/goblin"
)

func Test(t *testing.T) {
	g := Goblin(t)

	initRegEx()

	g.Describe("GAME INIT parse", func() {
		g.After(func() {
			games = make(map[string]*game)
		})

		g.It("Should return as checked and has no errors", func() {
			check, err := checkAndParseGameInit(`0:00 InitGame: \sv_floodProtect\1\sv_maxPing\0\sv_minPing\0\sv_maxRate\10000\sv_minRate\0\sv_hostname\Code Miner Server\g_gametype\0\sv_privateClients\2\sv_maxclients\16\sv_allowDownload\0\dmflags\0\fraglimit\20\timelimit\15\g_maxGameClients\0\capturelimit\8\version\ioq3 1.36 linux-x86_64 Apr 12 2009\protocol\68\mapname\q3dm17\gamename\baseq3\g_needpass\0`)

			g.Assert(check).Equal(true)
			g.Assert(err).Equal(nil)
		})
	})

	g.Describe("USER parse", func() {
		g.Before(func() {
			checkAndParseGameInit(`0:00 InitGame: \sv_floodProtect\1\sv_maxPing\0\sv_minPing\0\sv_maxRate\10000\sv_minRate\0\sv_hostname\Code Miner Server\g_gametype\0\sv_privateClients\2\sv_maxclients\16\sv_allowDownload\0\dmflags\0\fraglimit\20\timelimit\15\g_maxGameClients\0\capturelimit\8\version\ioq3 1.36 linux-x86_64 Apr 12 2009\protocol\68\mapname\q3dm17\gamename\baseq3\g_needpass\0`)
		})

		g.After(func() {
			games = make(map[string]*game)
		})

		g.It("Should return as checked and has no errors", func() {
			check, err := checkAndParseUser(`ClientUserinfoChanged: 2 n\Isgalamido\t\0\model\xian/default\hmodel\xian/default\g_redteam\\g_blueteam\\c1\4\c2\5\hc\100\w\0\l\0\tt\0\tl\0`)

			g.Assert(check).Equal(true)
			g.Assert(err).Equal(nil)
		})
	})

	g.Describe("KILL parse", func() {
		g.Before(func() {
			checkAndParseGameInit(`0:00 InitGame: \sv_floodProtect\1\sv_maxPing\0\sv_minPing\0\sv_maxRate\10000\sv_minRate\0\sv_hostname\Code Miner Server\g_gametype\0\sv_privateClients\2\sv_maxclients\16\sv_allowDownload\0\dmflags\0\fraglimit\20\timelimit\15\g_maxGameClients\0\capturelimit\8\version\ioq3 1.36 linux-x86_64 Apr 12 2009\protocol\68\mapname\q3dm17\gamename\baseq3\g_needpass\0`)
		})

		g.After(func() {
			games = make(map[string]*game)
		})

		g.It("Should return as checked and has no errors", func() {
			check, err := checkAndParseKill(`0:17 Kill: 1022 3 22: <world> killed Dono da Bola by MOD_TRIGGER_HURT`)

			g.Assert(check).Equal(true)
			g.Assert(err).Equal(nil)
		})
	})
}
