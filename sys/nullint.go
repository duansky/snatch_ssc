package sys

import "encoding/json"

type NullInt struct {
	int64
	Valid bool
}

func NewNullInt(i int64) *NullInt {
	return &NullInt{i, true}
}

func (this *NullInt) MarshalJSON() ([]byte, error) {
	if !this.Valid {
		return []byte(`""`), nil
	}

	return json.Marshal(this.int64)
}

func (this *NullInt) UnmarshalJSON(vals []byte) error {
	var v int64
	if err := json.Unmarshal(vals, &v); err != nil {
		return err
	}
	this.int64 = v
	this.Valid = true
	return nil
}

func (this *NullInt) Scan(value interface{}) error {
	if value == nil {
		this.int64, this.Valid = 0, false
		return nil
	}
	this.Valid = true
	return convertAssign(&this.int64, value)
}

func (this *NullInt) Value() int64 {
	if this.Valid {
		return this.int64
	}
	return 0
}

func (this *NullInt) SetValue(value int64) {
	this.int64 = value
	this.Valid = true
}
