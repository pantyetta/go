package main

import "fmt"

type Account struct {
	fitstName, lastName string
}

func (a *Account) ChangeName(firstName, lastName string) {
	a.fitstName = firstName
	a.lastName = lastName
}

type Employee struct {
	Account
	credits float64
}

func (e *Employee) AddCredits(credits float64) {
	e.credits += credits
}

func (e *Employee) RemoveCredits(credits float64) {
	e.credits -= credits
}

func (e Employee) CheckCredits() float64 {
	return e.credits
}

func (e Employee) String() string {
	return fmt.Sprintf("%s%s", e.fitstName, e.lastName)
}

func main() {
	employee := Employee{Account{"pan", "cho"}, 30}
	fmt.Println(employee)
	employee.ChangeName("cho", "pan")
	employee.AddCredits(100)
	employee.RemoveCredits(50)
	fmt.Println(employee.CheckCredits())
}
