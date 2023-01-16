package main

import "fmt"

func main() {
	studentsAge := make(map[string]int)
	studentsAge["John"] = 32
	studentsAge["bob"] = 31
	age, exist := studentsAge["christy"]
	if exist {
		fmt.Println("Chisty's age is", age)
	} else {
		fmt.Println("Christy's age couldn't be found.")
	}
}
