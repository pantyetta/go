package main

import "fmt"

func main() {
	studentsAge := make(map[string]int)
	studentsAge["John"] = 32
	studentsAge["bob"] = 31
	fmt.Println(studentsAge)
}
