package base

import (
	"database/sql/driver"
	"github.com/lib/pq"
)

type StringArray pq.StringArray

func (a *StringArray) Scan(src interface{}) error {
	receiver := (*pq.StringArray)(a)
	return receiver.Scan(src)
}

func (a StringArray) Value() (driver.Value, error) {
	return ((pq.StringArray)(a)).Value()
}
