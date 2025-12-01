package main

import (
	"fmt"
	"possum/final"
)

func main() {
	val, err := final.Run("input")
	if err != nil {
		panic(err)
	}
	fmt.Println("Value: ", val)
}
