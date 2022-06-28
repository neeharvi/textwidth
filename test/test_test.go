package textwidth

import (
	"testing"

	"github.com/neeharvi/textwidth"
)

// test cases copied from https://github.com/mattn/go-runewidth/raw/master/runewidth_test.go

var stringwidthtests = []struct {
	in    string
	out   int
	eaout int
}{
	{"■㈱の世界①", 10, 12},
	{"スター☆", 7, 8},
	{"つのだ☆HIRO", 11, 12},
}

func TestWidthString(t *testing.T) {
	for _, tt := range stringwidthtests {
		if out := textwidth.WidthString(tt.in); out != tt.out {
			t.Errorf("WidthString(%q) = %d, want %d", tt.in, out, tt.out)
		}
	}
}

func TestWidthStringAsian(t *testing.T) {
	for _, tt := range stringwidthtests {
		if out := textwidth.WidthStringCJK(tt.in, true); out != tt.eaout {
			t.Errorf("WidthStringAsian(%q) = %d, want %d", tt.in, out, tt.eaout)
		}
	}
}

var slicewidthtests = []struct {
	in    []byte
	out   int
	eaout int
}{
	{[]byte("■㈱の世界①"), 10, 12},
	{[]byte("スター☆"), 7, 8},
	{[]byte("つのだ☆HIRO"), 11, 12},
}

func TestWidthBytes(t *testing.T) {
	for _, tt := range slicewidthtests {
		if out := textwidth.WidthBytes(tt.in); out != tt.out {
			t.Errorf("WidthBytes(%q) = %d, want %d", tt.in, out, tt.out)
		}
	}
}

func TestWidthBytesAsian(t *testing.T) {
	for _, tt := range slicewidthtests {
		if out := textwidth.WidthBytesCJK(tt.in, true); out != tt.eaout {
			t.Errorf("WidthBytesAsian(%q) = %d, want %d", tt.in, out, tt.eaout)
		}
	}
}

var runewidthtests = []struct {
	in     rune
	out    int
	eaout  int
	nseout int
}{
	/* 00 */ {'世', 2, 2, 2},
	/* 01 */ {'界', 2, 2, 2},
	/* 02 */ {'ｾ', 1, 1, 1},
	/* 03 */ {'ｶ', 1, 1, 1},
	/* 04 */ {'ｲ', 1, 1, 1},
	/* 05 */ {'☆', 1, 2, 2}, // double width in ambiguous
	/* 06 */ {'☺', 1, 1, 2},
	/* 07 */ {'☻', 1, 1, 2},
	/* 08 */ {'♥', 1, 2, 2},
	/* 09 */ {'♦', 1, 1, 2},
	/* 10 */ {'♣', 1, 2, 2},
	/* 11 */ {'♠', 1, 2, 2},
	/* 12 */ {'♂', 1, 2, 2},
	/* 13 */ {'♀', 1, 2, 2},
	/* 14 */ {'♪', 1, 2, 2},
	/* 15 */ {'♫', 1, 1, 2},
	/* 16 */ {'☼', 1, 1, 2},
	/* 17 */ {'↕', 1, 2, 2},
	/* 18 */ {'‼', 1, 1, 2},
	/* 19 */ {'↔', 1, 2, 2},
	/* 20 */ {'\x00', 0, 0, 0},
	/* 21 */ {'\x01', 0, 0, 0},
	/* 22 */ {'\u0300', 0, 0, 0},
	/* 23 */ {'\u2028', 0, 0, 0},
	/* 24 */ {'\u2029', 0, 0, 0},
	/* 25 */ {'a', 1, 1, 1}, // ASCII classified as "na" (narrow)
	/* 26 */ {'⟦', 1, 1, 1}, // non-ASCII classified as "na" (narrow)
	/* 27 */ {'👁', 1, 1, 2},
}

func TestWidthRune(t *testing.T) {
	for i, tt := range runewidthtests {
		if out := textwidth.WidthRune(tt.in); out != tt.out {
			t.Errorf("case %d: WidthRune(%q) = %d, want %d", i, tt.in, out, tt.out)
		}
	}
}

func TestWidthRuneAsian(t *testing.T) {
	for i, tt := range runewidthtests {
		if out := textwidth.WidthRuneCJK(tt.in, true); out != tt.eaout {
			t.Errorf("case %d: WidthRuneAsian(%q) = %d, want %d", i, tt.in, out, tt.eaout)
		}
	}
}
