package i18n

import (
	"strings"

	"github.com/leekchan/accounting"
)

const (
	CURRENCY_EUR = "EUR"
	CURRENCY_NOK = "NOK"
	CURRENCY_USD = "USD"
	CURRENCY_MXN = "MXN"

	COUNTRY_ES = "ES"
	COUNTRY_NO = "NO"
	COUNTRY_EN = "EN"
	COUNTRY_MX = "MX"
)

type nodes_t map[string]interface{}
type textos_t map[string]string
type lang_t struct {
	textos textos_t
}

type I18n struct {
	ac       *accounting.Accounting
	currency string
	country  string
	lang     string

	fmtDate       string
	fmtTime       string
	fmtDateTime   string
	fmtDateTimeTz string

	langs map[string]*lang_t
}

func New() *I18n {

	f := new(I18n)
	f.langs = make(map[string]*lang_t)

	f.SetMoney(CURRENCY_EUR) // Default currency
	f.SetCountry(COUNTRY_ES) // Default country

	return f
}

func (f *I18n) SetLang(lang string) {
	f.lang = lang
}

func (f *I18n) SetMoney(currency string) {

	f.currency = strings.ToUpper(currency)

	switch f.currency {
	case CURRENCY_EUR:
		f.ac = accounting.NewAccounting("€", 2, ".", ",", "%v %s", "-%v %s", "")
	case CURRENCY_NOK:
		f.ac = accounting.NewAccounting("kr", 2, ".", ",", "%v %s", "-%v %s", "")
	case CURRENCY_MXN:
		f.ac = accounting.NewAccounting("$", 2, ",", ".", "%s %v", "-%s %v", "")
	default:
		f.ac = accounting.NewAccounting("$", 2, ",", ".", "%s %v", "-%s %v", "")
	}
}

func (f *I18n) SetCountry(country string) {
	f.country = strings.ToUpper(country)

	switch f.country {
	case COUNTRY_ES:
		f.fmtDate = "02/01/2006"
		f.fmtTime = "15:04:05"
		f.fmtDateTime = "02/01/06 15:04:05"
		f.fmtDateTimeTz = "02/01/06 15:04:05 -0700"
	case COUNTRY_NO:
		f.fmtDate = "02/01/2006"
		f.fmtTime = "15:04:05"
		f.fmtDateTime = "02/01/06 15:04:05"
		f.fmtDateTimeTz = "02/01/06 15:04:05 -0700"
	case COUNTRY_MX:
		f.fmtDate = "02/01/2006"
		f.fmtTime = "15:04:05"
		f.fmtDateTime = "02/01/06 15:04:05"
		f.fmtDateTimeTz = "02/01/06 15:04:05 -0700"

	default:
		f.fmtDate = "01/02/2006"
		f.fmtTime = "15:04:05"
		f.fmtDateTime = "01/02/06 15:04:05"
		f.fmtDateTimeTz = "01/02/06 15:04:05 -0700"
	}
}
