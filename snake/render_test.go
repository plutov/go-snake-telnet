package snake

import (
	"testing"
)

func BenchmarkRender(b *testing.B) {
	game := NewGame()

	for n := 0; n < b.N; n++ {
		game.Render()
	}
}
