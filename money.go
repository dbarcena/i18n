package i18n

import (
	"github.com/leekchan/accounting"
)

/*
Representa monedas en formato centavos o centimos como si fuera $ o €, siempre / 100.00
*/
func (f *I18n) Cent(val *int64) string {
	if val == nil {
		return ""
	} else {
		return f.ac.FormatMoneyFloat64(float64(*val) / 100.00)
	}
}

func (f *I18n) CentN(val *int64) string {
	if val == nil {
		return ""
	} else {

		return accounting.FormatNumberFloat64(
			float64(*val)/100.00,
			f.ac.Precision,
			f.ac.Thousand,
			f.ac.Decimal,
		)
	}
}

func (f *I18n) Money(val *float64) string {
	if val == nil {
		return ""
	} else {
		return f.ac.FormatMoneyFloat64(*val)
	}
}

func (f *I18n) MoneyN(val *float64) string {
	if val == nil {
		return ""
	} else {
		return accounting.FormatNumber(
			*val,
			f.ac.Precision,
			f.ac.Thousand,
			f.ac.Decimal,
		)
	}
}
