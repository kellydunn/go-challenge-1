package drum

import (
	"bytes"
)

type StepSequence struct {
	Steps []byte
}

func (s StepSequence) String() string {
	buf := bytes.NewBufferString("")

	for i, step := range s.Steps {
		if i%4 == 0 {
			buf.WriteString("|")
		}

		if bytes.Compare([]byte{step}, []byte{0}) == 0 {
			buf.WriteString("-")
		} else if bytes.Compare([]byte{step}, []byte{1}) == 0 {
			buf.WriteString("x")
		}
	}

	buf.WriteString("|")

	return buf.String()
}
