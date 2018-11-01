package drawer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wrutkowski/go1010/game"
)

func TestMergeBoardsHorizontally(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		inSeparator string
		in          []string
		want        string
	}{
		{" ", []string{"1\n1\n1", "2\n2\n2", "3\n3\n3"}, "1 2 3 \n1 2 3 \n1 2 3 "},
		{"", []string{"1\n1\n1", "\n\n2", "3\n3\n3"}, "13\n13\n123"},
		{" ", []string{"1\n1\n1", "2\n2", "3"}, "1 2 3 \n1 2 \n1 "},
		{"", []string{"", "2\n2", "3\n3\n3"}, "23\n23\n3"}}

	for _, c := range cases {
		assert.Equal(c.want, mergeBoardsHorizontally(c.inSeparator, c.in...))
	}
}

func TestWindowAround(t *testing.T) {
	assert := assert.New(t)
	window := ""

	content := "333     \n88888888\n1       \n        \n22      "

	window = windowAround("T", content, 8)
	assert.Equal("┌─ T ────┐\n│333     │\n│88888888│\n│1       │\n│        │\n│22      │\n└────────┘\n", window)

	window = windowAround("very long text", content, 8)
	assert.Equal("┌─ very long text ──┐\n│333                │\n│88888888           │\n│1                  │\n│                   │\n│22                 │\n└───────────────────┘\n", window)

	content = "\u00A9\u00A9\u00A9 \n\u250F   \n\u00A9\u00A9\u00A9\u00A9"
	window = windowAround("?", content, 4)
	assert.Equal("┌─ ? ──┐\n│©©©   │\n│┏     │\n│©©©©  │\n└──────┘\n", window)
}

func TestDrawBoard(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		in   [][]game.BoardElement
		want string
	}{
		{[][]game.BoardElement{{game.None, game.None}, {game.None, game.None}}, "  0 1  \n ┏━━━━┓\n0┃\x1b[40m \x1b[0m\x1b[40m \x1b[0m\x1b[40m \x1b[0m\x1b[40m \x1b[0m┃\n1┃\x1b[40m \x1b[0m\x1b[40m \x1b[0m\x1b[40m \x1b[0m\x1b[40m \x1b[0m┃\n ┗━━━━┛"}}

	for _, c := range cases {
		assert.Equal(c.want, drawBoard(c.in))
	}

}
