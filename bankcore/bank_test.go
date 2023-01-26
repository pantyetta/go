package bank

import "testing"

func TestAccount(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	if account.Name == "" {
		t.Error("Can't create an Account object")
	}
}
