package invalids

import (
	"strconv"
)

func GetFromRange(startStr, endStr string) []int {
	invalids := make([]int, 0)
	start, _ := strconv.Atoi(startStr)
	end, _ := strconv.Atoi(endStr)
	for num := start; num <= end; num += 1 {
		s := strconv.Itoa(num)
		ssize := len(s)
		if ssize%2 != 0 {
			continue
		}
		if s[:ssize/2] == s[ssize/2:] {
			invalids = append(invalids, num)
		}
	}
	return invalids
}
