package invalids

import (
	"strconv"
	"strings"
)

func GetRepeatingFromRange(startStr, endStr string) (invalids []int) {
	start, err := strconv.Atoi(startStr)
	if err != nil {
		panic(err)
	}
	end, err := strconv.Atoi(strings.TrimRight(endStr, "\n"))
	if err != nil {
		panic(err)
	}
	for i := start; i <= end; i += 1 {
		s := strconv.Itoa(i)
		ssize := len(s)
		if ssize%2 != 0 {
			continue
		}
		if s[:ssize/2] == s[ssize/2:] {
			invalids = append(invalids, i)
		}
	}
	return invalids
}
