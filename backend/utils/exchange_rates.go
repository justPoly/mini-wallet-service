package utils

import (
	"fmt"
)

var ExchangeRates = map[string]float64{
	"USD_NGN": 1550,
	"USD_EUR": 0.92,
	"USD_GBP": 0.79,
}

func Convert(amount float64, fromCurrency, toCurrency string) (float64, error) {

	if fromCurrency == toCurrency {
		return amount, nil
	}

	// Direct conversion
	key := fmt.Sprintf("%s_%s", fromCurrency, toCurrency)

	if rate, ok := ExchangeRates[key]; ok {
		return amount * rate, nil
	}

	// Reverse conversion
	reverseKey := fmt.Sprintf("%s_%s", toCurrency, fromCurrency)

	if rate, ok := ExchangeRates[reverseKey]; ok {
		return amount / rate, nil
	}

	return 0, fmt.Errorf("unsupported currency conversion from %s to %s", fromCurrency, toCurrency)
}