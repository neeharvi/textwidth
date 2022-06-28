// textwidth provides functions for getting the fixed-width width of unicode
// byte slices, runes and strings.
//
// Set the cjk flag to treat EastAsianAmbiguous as 2.
//
// https://en.wikipedia.org/wiki/Halfwidth_and_fullwidth_forms#In_Unicode
// https://unicode.org/reports/tr11/
// https://pkg.go.dev/golang.org/x/text/width
//
package textwidth

import (
	"unicode"
	"unsafe"

	"github.com/rivo/uniseg"
	"golang.org/x/text/width"
)

func WidthByteCJK(b byte, cjk bool) int {
	return WidthRuneCJK(rune(b), cjk)
}

func WidthRuneCJK(r rune, cjk bool) int {
	// TODO: better way?
	if !unicode.IsGraphic(r) || unicode.Is(unicode.Mn, r) {
		return 0
	}
	//table := [...]int{
	//	width.EastAsianWide: 2, width.EastAsianFullwidth: 2,
	//	width.EastAsianAmbiguous: 1,
	//	width.EastAsianNarrow:    1, width.EastAsianHalfwidth: 1, width.Neutral: 1,
	//}
	switch width.LookupRune(r).Kind() {
	case width.EastAsianWide, width.EastAsianFullwidth:
		return 2
	case width.EastAsianAmbiguous:
		if cjk {
			return 2
		} else {
			return 1
		}
	default:
		return 1
	}
}

func WidthGraphemeCJK(rs []rune, cjk bool) (n int) {
	// max width
	for _, r := range rs {
		w := WidthRuneCJK(r, cjk)
		if w > n {
			n = w
		}
	}
	return
}

func WidthBytesCJK(s []byte, cjk bool) (n int) {
	return WidthStringCJK(*(*string)(unsafe.Pointer(&s)), cjk)
}

func WidthStringCJK(s string, cjk bool) (n int) {
	// TODO: Is it necessary to iterate over graphemes here?
	//       mattn/go-runewidth does it like this.
	gs := uniseg.NewGraphemes(s)
	for gs.Next() {
		n += WidthGraphemeCJK(gs.Runes(), cjk)
	}
	return
}

func WidthByte(b byte) int       { return WidthByteCJK(b, false) }
func WidthRune(r rune) int       { return WidthRuneCJK(r, false) }
func WidthGrapheme(g []rune) int { return WidthGraphemeCJK(g, false) }
func WidthBytes(s []byte) int    { return WidthBytesCJK(s, false) }
func WidthString(s string) int   { return WidthStringCJK(s, false) }
