package sys

import (
	"encoding/json"
	"time"
)

const DATETIME_FMT = "2006-01-02 15:04:05"

type NullTime struct {
	time.Time
	Valid bool
}

func EmptyTime() *NullTime {
	return &NullTime{time.Time{}, false}
}

func NewNullTime(t time.Time) *NullTime {
	return &NullTime{t, !t.IsZero()}
}

func (this *NullTime) MarshalJSON() ([]byte, error) {
	if this.IsZero() {
		return []byte(`""`), nil
	}

	return json.Marshal(this.Format(DATETIME_FMT))
}

func (this *NullTime) UnmarshalJSON(vals []byte) error {
	var v time.Time
	if err := json.Unmarshal(vals, &v); err != nil {
		return err
	}
	this.Time = v
	this.Valid = true
	return nil
}

func (this *NullTime) Scan(value interface{}) error {
	if value == nil {
		this.Time, this.Valid = time.Time{}, false
		return nil
	}
	this.Valid = true
	return convertAssign(&this.Time, value)
}

//func (this *NullTime) Scan(value interface{}) error {
//	var t time.Time
//	if value == nil {
//		t, this.Valid = time.Time{}, false
//		return nil
//	}

//	if err := convertAssign(&t, value); err != nil {
//		this.Valid = false
//		return err
//	} else {
//		this.Time = t.In(time.FixedZone("", this.Offset))
//		this.Valid = true
//		return nil
//	}
//}

func (this *NullTime) Value() time.Time {
	if this.Valid {
		return this.Time
	}
	return time.Time{}
}

func (this *NullTime) SetValue(value time.Time) {
	this.Time = value
	this.Valid = true
}
