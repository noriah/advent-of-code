package day01

import "fmt"

func NewRunner() func(data []byte) {
	return func(data []byte) {
		fmt.Printf("Problem 1: %d\n", Problem1(data))
		fmt.Printf("Problem 2: %d\n", Problem2(data))
	}
}
