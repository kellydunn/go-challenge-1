package drum

import (
	"bytes"
	"testing"
)

var TEST_VERSION []byte = []byte{
	0x30, 0x2e, 0x38, 0x30,
	0x38, 0x2d, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x00,
	0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00,
}

var TEST_TEMPO []byte = []byte{
	0x00, 0x00, 0xf0, 0x42,
	0x00, 0x00, 0x00, 0x00,
}

func TestReadPatternVersion(t *testing.T) {
	p := &Pattern{}
	expected := "0.808-alpha"
	reader := bytes.NewReader(TEST_VERSION)

	read, err := readPatternVersion(reader, p)
	if err != nil {
		t.Errorf("Error reading track name %v", err)
	}

	if p.Version != expected {
		t.Errorf("Mismatched pattern version.  Expected: %d.  Actual: %d", expected, p.Version)
	}

	if read != VERSION_SIZE {
		t.Errorf("Mismatched bytes read. Expected %d.  Actual: %d", VERSION_SIZE, read)
	}
}

func TestReadPatternTempo(t *testing.T) {
	p := &Pattern{}
	expected := float32(120.0)
	reader := bytes.NewReader(TEST_TEMPO)

	read, err := readPatternTempo(reader, p)
	if err != nil {
		t.Errorf("Error reading track name %v", err)
	}

	if p.Tempo != expected {
		t.Errorf("Mismatched pattern tempo.  Expected: %d.  Actual: %d", expected, p.Tempo)
	}

	if read != TEMPO_SIZE {
		t.Errorf("Mismatched bytes read. Expected %d.  Actual: %d", TEMPO_SIZE, read)
	}
}
