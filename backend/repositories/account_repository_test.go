package repositories

import (
	"testing"

	"github.com/justPoly/mini-wallet-service/backend/models"
)

func TestCreateAccountModel(t *testing.T) {

	account := models.Account{
		Name: "John Doe",
		Currency: "USD",
	}

	if account.Name != "John Doe" {
		t.Error("Name not assigned")
	}

	if account.Currency != "USD" {
		t.Error("Currency incorrect")
	}
}