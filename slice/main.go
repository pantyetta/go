package main

import "fmt"

func main() {
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("%d\t%d\tcap=%d\t%v\n", i, len(numbers), cap(numbers), numbers)
	}
}
