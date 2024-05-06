package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Date struct {
	time.Time
}

const layout = "2006-01-02"

func (d *Date) UnmarshalJSON(b []byte) (err error) {
	s := string(b)
	s = s[1 : len(s)-1] // Remove quotes
	d.Time, err = time.Parse(layout, s)
	return
}

// MarshalJSON implements the json.Marshaler interface
func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", d.Format(layout))), nil
}

// Value implements the driver.Valuer interface
func (d Date) Value() (driver.Value, error) {
	return d.Format(layout), nil
}

// Scan implements the sql.Scanner interface
func (d *Date) Scan(value interface{}) error {
	if value == nil {
		*d = Date{Time: time.Time{}}
		return nil
	}
	if t, ok := value.(time.Time); ok {
		*d = Date{Time: t}
		return nil
	}
	return fmt.Errorf("can't convert %T to Date", value)
}
