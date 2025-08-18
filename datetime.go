package i18n

import (
	"time"
)

const (
	PGSQL = "2006-01-02T15:04:05.000000-07:00"
	MSSQL = "2006-01-02T15:04:05.000"
)

func (f *I18n) Date(val *string) string {
	if val == nil {
		return ""
	} else {

		if t, err := f.Parse(*val); err == nil {
			return t.Format(f.fmtDate)
		} else {
			return "ERROR"
		}
	}
}

func (f *I18n) Time(val *string) string {
	if val == nil {
		return ""
	} else {

		if t, err := f.Parse(*val); err == nil {
			return t.Format(f.fmtTime)
		} else {
			return "ERROR"
		}
	}
}

func (f *I18n) DateTime(val *string) string {
	if val == nil {
		return ""
	} else {
		if t, err := f.Parse(*val); err == nil {
			return t.Format(f.fmtDateTime)
		} else {
			return "ERROR"
		}
	}
}

func (f *I18n) DateTimeTz(val *string) string {
	if val == nil {
		return ""
	} else {

		if t, err := f.Parse(*val); err == nil {
			return t.Format(f.fmtDateTimeTz)
		} else {
			return "ERROR"
		}
	}
}

func (f *I18n) Now() string {
	return time.Now().Format(time.RFC3339Nano)
}

func (f *I18n) Parse(val string) (time.Time, error) {

	if t, err := time.Parse(PGSQL, val); err == nil {
		return t, err
	} else if t, err := time.Parse(MSSQL, val); err == nil {
		return t, err
	} else if t, err := time.Parse(time.RFC3339Nano, val); err == nil {
		return t, err
	} else {
		return time.Time{}, err
	}
}
