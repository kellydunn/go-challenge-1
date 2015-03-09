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

		if step == byte(0) {
			buf.WriteString("-")
		} else if step == byte(1) {
			buf.WriteString("x")
		}
	}

	buf.WriteString("|")

	return buf.String()
}
