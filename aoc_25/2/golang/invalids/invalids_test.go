package invalids_test

import (
	"possum/invalids"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFromRange(t *testing.T) {
	type testCase struct {
		From     string
		To       string
		Invalids []int
	}
	total := 0
	testCases := []testCase{
		{"11", "22", []int{11, 22}},
		{"95", "115", []int{99}},
		{"998", "1012", []int{1010}},
		{"1188511880", "1188511890", []int{1188511885}},
		{"222220", "222224", []int{222222}},
		{"1698522", "1698528", []int{}},
		{"446443", "446449", []int{446446}},
		{"38593856", "38593862", []int{38593859}},
		{"565653", "565659", []int{}},
		{"824824821", "824824827", []int{}},
		{"2121212118", "2121212124", []int{}},
	}
	for _, c := range testCases {
		out := invalids.GetFromRange(c.From, c.To)
		assert.Equal(t, out, c.Invalids)
		for _, r := range out {
			total += r
		}
	}
	assert.Equal(t, total, 1227775554)
}
