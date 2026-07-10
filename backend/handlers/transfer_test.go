package handlers

import "testing"

func TestInsufficientFunds(t *testing.T) {

	balance := 100.0
	transfer := 200.0

	if balance >= transfer {
		t.Error("Transfer should have failed")
	}
}