package main

import (
	"errors"
	"fmt"
)

type Account struct {
	fitstName, lastName string
}

type Employee struct {
	Account
	Credits float64
}

func (a *Account) ChangeName(firstName, lastName string) {
	a.fitstName = firstName
	a.lastName = lastName
}

func CreateEmployee(firstName, lastName string, credits float64) (*Employee, error) {
	return &Employee{Account{firstName, lastName}, credits}, nil
}

func (e *Employee) AddCredits(credits float64) error {
	if credits < 0.0 {
		return errors.New("Invalid credits amount.")
	}
	e.Credits += credits
	return nil
}

func (e *Employee) RemoveCredits(credits float64) error {
	if credits < 0.0 {
		return errors.New("Invalid credits amount.")
	}
	if e.Credits < credits {
		return errors.New("You can't remove credits than the account has.")
	}
	e.Credits -= credits
	return nil
}

func (e Employee) CheckCredits() float64 {
	return e.Credits
}

func (e Employee) String() string {
	return fmt.Sprintf("Name: %s %s\nCredits: %.2f\n", e.fitstName, e.lastName, e.Credits)
}

func main() {
	pan, _ := CreateEmployee("pan", "cho", 300)
	fmt.Println(pan.CheckCredits())
	err := pan.AddCredits(500)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = pan.RemoveCredits(2500)
	if err != nil {
		fmt.Println("Error:", err)
	}
	pan.ChangeName("pan", "tyetta")

	fmt.Println(pan)
}
