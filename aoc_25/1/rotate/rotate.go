package rotate

import "fmt"

func Rotate(pos, dir int) (int, int) {
	movesViaZero := 0
	fmt.Println("pos:", pos, "dir:", dir)
	if dir < 0 {
		abs_dir := 0 - dir
		movesViaZero += abs_dir / 100
		abs_dir %= 100
		dir = 0 - abs_dir
		if abs_dir > pos {
			// trim it
			if pos > 0 {
				movesViaZero += 1
			}
			dir = pos + dir
			pos = 100
		}
	}
	pos = pos + dir
	if pos == 0 {
		movesViaZero += 1
	} else {
		movesViaZero += pos / 100
	}
	pos %= 100
	return pos, movesViaZero
}
