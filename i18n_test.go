package i18n

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrency(t *testing.T) {

	i18n := New()
	i18n.SetMoney(CURRENCY_EUR)

	var f64 float64

	// EUROS
	assert.Equal(t, i18n.Cent(nil), "", "Expected empty string for nil value")

	f64 = 1234.56
	assert.Equal(t, i18n.Money(&f64), "1.234,56 €", "Positivo €")

	f64 = -1234.56
	assert.Equal(t, i18n.Money(&f64), "-1.234,56 €", "Negativo €")

	// NORWEGIAN KRONER
	i18n.SetMoney(CURRENCY_NOK)

	assert.Equal(t, i18n.Cent(nil), "", "Expected empty string for nil value")

	f64 = 1234.56
	assert.Equal(t, i18n.Money(&f64), "1.234,56 kr", "Positivo kr")

	f64 = -1234.56
	assert.Equal(t, i18n.Money(&f64), "-1.234,56 kr", "Negativo kr")

	// dolares
	i18n.SetMoney(CURRENCY_USD)

	assert.Equal(t, i18n.Cent(nil), "", "Expected empty string for nil value")

	f64 = 1234.56
	assert.Equal(t, i18n.Money(&f64), "$ 1,234.56", "Positivo $")

	f64 = -1234.56
	assert.Equal(t, i18n.Money(&f64), "-$ 1,234.56", "Negativo $")

	// Mx
	i18n.SetMoney(CURRENCY_MXN)

	assert.Equal(t, i18n.Cent(nil), "", "Expected empty string for nil value")

	f64 = 1234.56
	assert.Equal(t, i18n.Money(&f64), "$ 1,234.56", "Positivo $")

	f64 = -1234.56
	assert.Equal(t, i18n.Money(&f64), "-$ 1,234.56", "Negativo $")
}
