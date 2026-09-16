package pretok_test

import (
	"errors"
	"slices"
	"strings"
	"testing"

	"github.com/gosimd/pretok"
)

func TestP01BufferedContracts(t *testing.T) {
	re := compiledPattern(t, patternMarks)
	for _, size := range []int{0, 1, 15, 16, 63, 64, 254, 255, 256, 257} {
		text := strings.Repeat("a", size)
		got, err := re.Split(text, nil)
		if err != nil || !slices.Equal(got, collect(t, re, text)) {
			t.Fatal("dispatch threshold changed matches", size)
		}
		if _, err := re.Split(text+"\xff", nil); !errors.Is(err, pretok.ErrInvalidUTF8) {
			t.Fatal("dispatch threshold skipped validation", size)
		}
	}
	for _, length := range []int{255, 256, 257, 1023, 1024, 1025, 8193} {
		text := strings.Repeat("a", length) + "'re\r\n" + strings.Repeat("\u2003", length) + "X 中😀１２\ufffd"
		want := collect(t, re, text)
		back := make([]pretok.Span, 1, len(want)+1)
		back[0] = pretok.Span{Start: 91, End: 92}
		got, err := re.Split(text, back)
		if err != nil || &got[0] != &back[0] || got[0] != (pretok.Span{Start: 91, End: 92}) || !slices.Equal(got[1:], want) {
			t.Fatal("append/reuse differs", length, err)
		}
		for _, bad := range []string{"\xff", "\xc0\xaf", "\xed\xa0\x80", "\xf4\x90\x80\x80", "\xe2", "\xf0\x9f"} {
			for _, input := range []string{bad + text, text + bad, text + bad + text} {
				backing := make([]pretok.Span, 64)
				for i := range backing {
					backing[i] = pretok.Span{Start: 91, End: 92}
				}
				original := slices.Clone(backing)
				dst, err := re.Split(input, backing[:1])
				if !errors.Is(err, pretok.ErrInvalidUTF8) || len(dst) != 1 || &dst[0] != &backing[0] || !slices.Equal(backing, original) {
					t.Fatal("invalid input modified destination", length)
				}
				c, err := re.Cursor(input)
				if !errors.Is(err, pretok.ErrInvalidUTF8) {
					t.Fatal("cursor failed to validate entire input")
				}
				if _, ok := c.Next(); ok {
					t.Fatal("invalid cursor yielded a match")
				}
			}
		}
	}
	testConcurrentReuse(t, re, strings.Repeat("we're a\u0301 中😀\r\n ", 256))
	// Public cursor copies stay independent across multiple would-be rune windows.
	text := strings.Repeat("we're a\u0301 中😀\r\n ", 1024)
	c, _ := re.Cursor(text)
	c.Next()
	copy := c
	for i := 0; i < 1025; i++ {
		c.Next()
	}
	got := []pretok.Span{}
	for s, ok := copy.Next(); ok; s, ok = copy.Next() {
		got = append(got, s)
	}
	want := collect(t, re, text)
	if !slices.Equal(got, want[1:]) {
		t.Fatal("copy shares cursor progress")
	}
}
