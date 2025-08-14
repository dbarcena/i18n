package i18n

import "github.com/leekchan/accounting"

const (
	CURRENCY_EUR = "EUR"
	CURRENCY_NOK = "NOK"
	CURRENCY_USD = "USD"
	CURRENCY_MXN = "MXN"

	COUNTRY_ES = "ES"
	COUNTRY_NO = "NO"
	COUNTRY_EN = "EN"
)

type I18n struct {
	ac *accounting.Accounting
}

func New() *I18n {

	i18n := new(I18n)

	return i18n
}

func (i18n *I18n) SetMoney(currency string) {

	switch currency {
	case CURRENCY_EUR:
		i18n.ac = accounting.NewAccounting("€", 2, ".", ",", "%v %s", "-%v %s", "")
	case CURRENCY_NOK:
		i18n.ac = accounting.NewAccounting("kr", 2, ".", ",", "%v %s", "-%v %s", "")
	case CURRENCY_MXN:
	default:
		i18n.ac = accounting.NewAccounting("$", 2, ",", ".", "%s %v", "-%s %v", "")
	}
}
