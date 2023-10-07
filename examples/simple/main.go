package main

import (
	"fmt"

	"github.com/danielgatis/go-iterator"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	iter := iterator.NewIterator(numbers)

	for iter.HasNext() {
		value := iter.GetNextOrDefault(0)
		fmt.Println(value)
	}
}
