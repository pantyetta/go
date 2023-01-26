package main

import (
	"testing"

	bank "github.com/pantyetta/bankcore"
)

func TestCreateAccount(t *testing.T) {
	accounts[0] = &bank.Account{
		Customer: bank.Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number: 1001,
	}

	if accounts[0].Name != "John" {
		t.Error("Cna't create an Account object on map")
	}
}
