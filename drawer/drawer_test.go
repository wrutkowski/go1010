package drawer

import (
	"testing"

	"github.com/wrutkowski/go1010/game"
)

func TestMergeBoardsHorizontally(t *testing.T) {

	cases := []struct {
		inSeparator string
		in          []string
		want        string
	}{
		{" ", []string{"1\n1\n1", "2\n2\n2", "3\n3\n3"}, "1 2 3 \n1 2 3 \n1 2 3 \n"},
		{"", []string{"1\n1\n1", "\n\n2", "3\n3\n3"}, "13\n13\n123\n"},
		{" ", []string{"1\n1\n1", "2\n2", "3"}, "1 2 3 \n1 2 \n1 \n"},
		{"", []string{"", "2\n2", "3\n3\n3"}, "23\n23\n3\n"}}

	for _, c := range cases {
		got := mergeBoardsHorizontally(c.inSeparator, c.in...)
		if got != c.want {
			t.Errorf("mergeBoardsHorizontally(%q, %q) == %q, want %q", c.inSeparator, c.in, got, c.want)
		}
	}
}

func TestDrawBoard(t *testing.T) {

	cases := []struct {
		in   [][]game.BoardElement
		want string
	}{
		{[][]game.BoardElement{{game.None, game.None}, {game.None, game.None}}, "┏━━━━┓\n┃\x1b[40m \x1b[0m\x1b[40m \x1b[0m\x1b[40m \x1b[0m\x1b[40m \x1b[0m┃\n┃\x1b[40m \x1b[0m\x1b[40m \x1b[0m\x1b[40m \x1b[0m\x1b[40m \x1b[0m┃\n┗━━━━┛\n"}}

	for _, c := range cases {
		got := drawBoard(c.in)
		if got != c.want {
			t.Errorf("drawBoard(%q) == %q, want %q", c.in, got, c.want)
		}
	}

}
