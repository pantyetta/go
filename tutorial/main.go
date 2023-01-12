package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, _ := strconv.Atoi("-42") // ascii to int
	s := strconv.Itoa(-42)      // int to ascii
	fmt.Println(i, s)
}
