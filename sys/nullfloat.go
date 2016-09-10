package sys

import "encoding/json"

type NullFloat struct {
	float64
	Valid bool
}

func NewNullFloat(f float64) *NullFloat {
	return &NullFloat{f, true}
}

func (this *NullFloat) MarshalJSON() ([]byte, error) {
	if !this.Valid {
		return []byte(`""`), nil
	}

	return json.Marshal(this.float64)
}

func (this *NullFloat) UnmarshalJSON(vals []byte) error {
	var v float64
	if err := json.Unmarshal(vals, &v); err != nil {
		return err
	}
	this.float64 = v
	this.Valid = true
	return nil
}

func (this *NullFloat) Scan(value interface{}) error {
	if value == nil {
		this.float64, this.Valid = 0, false
		return nil
	}
	this.Valid = true
	return convertAssign(&this.float64, value)
}

func (this *NullFloat) Value() float64 {
	if this.Valid {
		return this.float64
	}
	return 0.00
}

func (this *NullFloat) SetValue(value float64) {
	this.float64 = value
	this.Valid = true
}
