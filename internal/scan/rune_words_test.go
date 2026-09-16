package scan

import (
	"math/rand"
	"strings"
	"testing"
	"unicode/utf8"
)

func assertRuneWords(t testing.TB, src string, size int) {
	t.Helper()
	// Guards catch writes outside the borrowed scratch, including refill copies.
	back := make([]rune, size+2)
	back[0] = 0x123456
	back[size+1] = 0x654321
	c := NewQwenMarksCursor(src, back[1:size+1])
	scalar := qwen(true, 1)
	pos := 0
	for start, end, ok := c.Next(); ok; start, end, ok = c.Next() {
		want := scalar.Match(src, pos)
		if start != pos || end != want || end <= start || end > len(src) {
			t.Fatalf("capacity=%d byte=%d got=[%d,%d) want=[%d,%d) input bytes=%d", size, pos, start, end, pos, want, len(src))
		}
		pos = end
	}
	if pos != len(src) {
		t.Fatalf("stopped at %d of %d", pos, len(src))
	}
	if a, b, ok := c.Next(); ok || a != 0 || b != 0 {
		t.Fatal("cursor not exhausted")
	}
	if back[0] != 0x123456 || back[size+1] != 0x654321 {
		t.Fatal("scratch guard overwritten")
	}
}
func TestRuneWordsBoundaries(t *testing.T) {
	cases := []string{"", "a", "\ufffd", "we're I'LL 'ſ 'RE 'Ve 'Ll", "!!'s", "a\u0301b", "\r\n\tA", "\u2003\u2003X", "１２²Ⅳ1"}
	for _, n := range []int{3, 4, 5, 7, 15, 16, 17, 1023, 1024, 1025, 2047, 2048, 2049, 4095, 4096, 4097, 8191, 8192, 8193, 65537} {
		for _, token := range []string{"a", "Ж", "中", "😀", "\u0301", " ", "\u2003", "\r\n"} {
			s := strings.Repeat(token, n)
			cases = append(cases, s, s+"X", "x"+s+"'re", "\n"+s+"X", "  \n"+s+"  !")
		}
	}
	rng := rand.New(rand.NewSource(20260915))
	alphabet := []rune("aAſ'stmrelv\r\n\t\u2003\u0301中١12²!😀\ufffd")
	for i := 0; i < 3000; i++ {
		r := make([]rune, rng.Intn(200))
		for j := range r {
			r[j] = alphabet[rng.Intn(len(alphabet))]
		}
		cases = append(cases, string(r))
	}

	for _, size := range []int{3, 4, 7, 16, 1024, 2048, 4096, 8192} {
		for _, src := range cases {
			assertRuneWords(t, src, size)
		}
	}
}
func FuzzRuneWords(f *testing.F) {
	for _, s := range []string{"", "we're 'ſ", "a\u0301b", "\r\n\u2003X", strings.Repeat("A", 1023) + "'re", strings.Repeat(" ", 1025) + "X", strings.Repeat("中😀", 1024), "\xff"} {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, src string) {
		if !utf8.ValidString(src) {
			return
		}
		for _, size := range []int{3, 16, 1024} {
			assertRuneWords(t, src, size)
		}
	})
}
