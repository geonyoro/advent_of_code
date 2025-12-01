package rotate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	type testCase struct {
		currentPos   int
		dir          int
		expectedPos  int
		movesViaZero int
	}
	pos := 50
	movesViaZero := 0
	newPos := 0
	inputs := []testCase{
		{50, -1000, 50, 10},
		{50, 1000, 50, 10},
		{50, -68, 82, 1},
		{82, -30, 52, 0},
		{52, 48, 0, 1},
		{0, -5, 95, 0},
		{95, 60, 55, 1},
		{55, -55, 0, 1},
		{0, -1, 99, 0},
		{99, -99, 0, 1},
		{0, 14, 14, 0},
		{14, -82, 32, 1},
	}
	for _, inp := range inputs {
		newPos, movesViaZero = Rotate(pos, inp.dir)
		actual := testCase{
			pos, inp.dir, newPos, movesViaZero,
		}
		ok := assert.Equal(t, inp, actual)
		if !ok {
			break
		}
		pos = newPos
	}
}
