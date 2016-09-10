package sys

import "encoding/json"

type NullString struct {
	string
	Valid bool
}

func EmptyString() *NullString {
	return &NullString{"", false}
}

func NewNullString(source string) *NullString {
	return &NullString{source, true}
}

func (this *NullString) MarshalJSON() ([]byte, error) {

	return json.Marshal(this.string)
}

func (this *NullString) UnmarshalJSON(vals []byte) error {
	var v string
	if err := json.Unmarshal(vals, &v); err != nil {
		return err
	}
	this.string = v
	this.Valid = true
	return nil
}

func (this *NullString) Scan(value interface{}) error {
	if value == nil {
		this.string, this.Valid = "", false
		return nil
	}
	this.Valid = true
	return convertAssign(&this.string, value)
}

func (this *NullString) String() string {
	if this != nil && this.Valid {
		return this.string
	}
	return ""
}

func (this *NullString) SetValue(value string) {
	this.string = value
	this.Valid = true
}
