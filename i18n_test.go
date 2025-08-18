package i18n

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test_currency_i64 struct {
	currency string
	importe  int64
	result   string
	resultN  string
}

type test_currency_f64 struct {
	currency string
	importe  float64
	result   string
	resultN  string
}

func TestCurrency(t *testing.T) {

	// -------------------------------------------
	list_i64 := []test_currency_i64{
		{CURRENCY_EUR, 123456, "1.234,56 €", "1.234,56"},
		{CURRENCY_NOK, 123456, "1.234,56 kr", "1.234,56"},
		{CURRENCY_USD, 123456, "$ 1,234.56", "1,234.56"},
		{CURRENCY_MXN, 123456, "$ 1,234.56", "1,234.56"},
	}

	// -------------------------------------------
	list_f64 := []test_currency_f64{
		{CURRENCY_EUR, 1234.56, "1.234,56 €", "1.234,56"},
		{CURRENCY_NOK, 1234.56, "1.234,56 kr", "1.234,56"},
		{CURRENCY_USD, 1234.56, "$ 1,234.56", "1,234.56"},
		{CURRENCY_MXN, 1234.56, "$ 1,234.56", "1,234.56"},
	}

	// -------------------------------------------
	// -------------------------------------------
	i18n := New()
	var last_currency string

	// -------------------------------------------
	i18n.SetMoney(CURRENCY_EUR)
	last_currency = CURRENCY_EUR

	for _, test := range list_i64 {
		if test.currency != last_currency {
			i18n.SetMoney(test.currency)
			last_currency = test.currency
		}
		assert.Equal(t, i18n.Cent(&test.importe), test.result, "Cent de "+test.currency)
		assert.Equal(t, i18n.CentN(&test.importe), test.resultN, "CentN de "+test.currency)
	}

	// -------------------------------------------
	i18n.SetMoney(CURRENCY_EUR)
	last_currency = CURRENCY_EUR

	for _, test := range list_f64 {
		if test.currency != last_currency {
			i18n.SetMoney(test.currency)
			last_currency = test.currency
		}
		assert.Equal(t, i18n.Money(&test.importe), test.result, "Money de "+test.currency)
		assert.Equal(t, i18n.MoneyN(&test.importe), test.resultN, "MoneyN de "+test.currency)
	}
}

func TestLang(t *testing.T) {

	f := New()
	if err := f.Langs("./testdata/languages"); err != nil {
		t.Fatalf("Error loading languages: %v", err)
	}

	f.SetLang("es-ES")

	assert.Equal(t, f.T("hola"), "Hola", "Translation for 'hola'")

	f.SetLang("en-EN")
	assert.Equal(t, f.T("hola"), "Hello", "Translation for 'hola'")

	f.SetLang("en-XX") // Non-existent language
	assert.Equal(t, f.T("hola"), "en-XX.hola", "ERROR en-XX not found")
}
