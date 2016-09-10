package sys

import (
	"bytes"
	"fmt"
)

type StringBuilder struct {
	buffer bytes.Buffer
}

func NewStringBuilder() *StringBuilder {
	return new(StringBuilder)
}

func (this *StringBuilder) AppendLine(line string) *StringBuilder {
	this.buffer.WriteString(fmt.Sprintln(line))
	return this
}

func (this *StringBuilder) AppendFormat(format string, a ...interface{}) *StringBuilder {
	this.AppendLine(fmt.Sprintf(format, a...))
	return this
}

func (this *StringBuilder) HasValue() bool {
	return this.buffer.Len() > 0
}

func (this *StringBuilder) String() string {
	return this.buffer.String()
}
