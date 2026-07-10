package utils

import "testing"

func TestConvertUSDToNGN(t *testing.T) {

	result, err := Convert(100, "USD", "NGN")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := 155000.0

	if result != expected {
		t.Errorf("Expected %.2f, got %.2f", expected, result)
	}
}

func TestUnsupportedCurrency(t *testing.T) {

	_, err := Convert(
		100,
		"USD",
		"JPY",
	)

	if err == nil {
		t.Fatal("expected error for unsupported currency")
	}
}