package drum

import (
	"bytes"
	"log"
)

type StepSequence struct {
	Steps []byte
}

func (s StepSequence) String() string {
	buf := bytes.NewBufferString("")

	for i, step := range s.Steps {
		if i%4 == 0 {
			_, err := buf.WriteString("|")
			if err != nil {
				log.Printf("Error writing to buffer: %v", err)
			}
		}

		if step == byte(0) {
			_, err := buf.WriteString("-")
			if err != nil {
				log.Printf("Error writing to buffer: %v", err)
			}
		} else if step == byte(1) {
			_, err := buf.WriteString("x")
			if err != nil {
				log.Printf("Error writing to buffer: %v", err)
			}
		}
	}

	_, err := buf.WriteString("|")
	if err != nil {
		log.Printf("Error writing to buffer: %v", err)
	}

	return buf.String()
}
