package sys

import (
	"encoding/json"
	"time"
)

const DATE_FMT = "2006-01-02"

type NullDate struct {
	time.Time
	Valid bool
}

func EmptyDate() *NullDate {
	return &NullDate{time.Time{}, false}
}

func (this *NullDate) MarshalJSON() ([]byte, error) {
	if this.IsZero() {
		return []byte(`""`), nil
	}

	return json.Marshal(this.Format(DATE_FMT))
}

func (this *NullDate) UnmarshalJSON(vals []byte) error {
	var v time.Time
	if err := json.Unmarshal(vals, &v); err != nil {
		return err
	}
	this.Time = v
	this.Valid = true
	return nil
}

func (this *NullDate) Scan(value interface{}) error {
	if value == nil {
		this.Time, this.Valid = time.Time{}, false
		return nil
	}
	this.Valid = true
	return convertAssign(&this.Time, value)
}

func (this *NullDate) Value() time.Time {
	if this.Valid {
		return this.Time
	}
	return time.Time{}
}

func (this *NullDate) SetValue(value time.Time) {
	this.Time = value
	this.Valid = true
}
