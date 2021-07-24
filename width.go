// textwidth provides functions for getting the fixed-width width of unicode
// byte slices, runes and strings.
//
// https://en.wikipedia.org/wiki/Halfwidth_and_fullwidth_forms#In_Unicode
//
package textwidth

import (
	"unicode"
	"unsafe"

	"golang.org/x/text/width"
)

var table = [...]int{
	width.EastAsianWide: 2, width.EastAsianFullwidth: 2,
	width.EastAsianNarrow: 1, width.EastAsianHalfwidth: 1, width.EastAsianAmbiguous: 1, width.Neutral: 1,
}

func WidthByte(b byte) int {
	return WidthRune(rune(b))
}

func WidthRune(r rune) int {
	if unicode.Is(unicode.Mn, r) || !unicode.IsGraphic(r) {
		return 0
	} else {
		return table[width.LookupRune(r).Kind()]
	}
}

func WidthBytes(s []byte) (n int) {
	return WidthString(*(*string)(unsafe.Pointer(&s)))
}

func WidthRunes(s []rune) (n int) {
	for _, r := range s {
		if unicode.Is(unicode.Mn, r) || !unicode.IsGraphic(r) {
			// no-op //
		} else {
			n += table[width.LookupRune(r).Kind()]
		}
	}
	return n
}

func WidthString(s string) (n int) {
	for _, r := range s {
		if unicode.Is(unicode.Mn, r) || !unicode.IsGraphic(r) {
			// no-op //
		} else {
			n += table[width.LookupRune(r).Kind()]
		}
	}
	return n
}
