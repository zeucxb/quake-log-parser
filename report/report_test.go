package report

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestReportSanitalizeGameName(t *testing.T) {
	g := Goblin(t)

	g.Describe("SanitalizeGameName", func() {
		g.It("Should return game_1 when receive 1", func() {
			sanitalizedGameName := sanitalizeGameName("1")

			g.Assert(sanitalizedGameName).Equal("game_1")
		})

		g.It("Should return game_1 when receive game_1", func() {
			sanitalizedGameName := sanitalizeGameName("game_1")

			g.Assert(sanitalizedGameName).Equal("game_1")
		})
	})
}
