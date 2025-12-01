package final

import (
	"os"
	"possum/rotate"
	"strconv"
	"strings"
)

func Run(filename string) (int, error) {
	inputB, err := os.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	input := string(inputB)
	stringIter := strings.SplitSeq(input, "\n")
	pos := 50
	atZero := 0
	var movesViaZero int
	for line := range stringIter {
		if len(line) == 0 {
			break
		}
		dir := line[0]
		rest := line[1:]
		num, err := strconv.Atoi(rest)
		if err != nil {
			return atZero, err
		}
		if dir == 'L' {
			num = 0 - num
		}
		pos, movesViaZero = rotate.Rotate(pos, num)
		atZero += movesViaZero
	}
	return atZero, nil
}
