package textwidth

import (
	"testing"

	"github.com/mattn/go-runewidth"
	"github.com/neeharvi/textwidth"
)

func BenchmarkStringWidth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		textwidth.WidthString(stringwidthtests[i%len(stringwidthtests)].in)
	}
}

func BenchmarkStringWidthOriginal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runewidth.StringWidth(stringwidthtests[i%len(stringwidthtests)].in)
	}
}

func BenchmarkRuneWidth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		textwidth.WidthRune(runewidthtests[i%len(runewidthtests)].in)
	}
}

func BenchmarkRuneWidthOriginal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runewidth.RuneWidth(runewidthtests[i%len(runewidthtests)].in)
	}
}
