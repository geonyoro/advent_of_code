package main

import (
	"fmt"
	"os"
	"possum/invalids"
	"strings"
)

func main() {
	inputBytes, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}
	input := string(inputBytes)
	stringIter := strings.SplitSeq(input, ",")
	total := 0
	for part := range stringIter {
		parts := strings.Split(part, "-")
		from := parts[0]
		to := parts[1]
		repeating := invalids.GetFromRange(from, to)
		fmt.Println(part, repeating)
		for _, r := range repeating {
			total += r
		}
	}
	fmt.Println("total:", total)
}
