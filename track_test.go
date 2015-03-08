package drum

import (
	"bytes"
	"testing"
)

var TEST_ID []byte = []byte{ 0x01 }

var TEST_NAME []byte = []byte{
	0x04, 0x6b, 0x69, 0x63, 0x6b,
}

var TEST_STEP_SEQUENCE []byte = []byte{
	0x01, 0x00, 0x00, 0x00,
	0x01, 0x01, 0x00, 0x00,
	0x01, 0x01, 0x01, 0x00,
	0x01, 0x01, 0x01, 0x01,
}

func TestReadTrackId(t *testing.T) {
	track := &Track{}
	expected := uint8(1)
	reader := bytes.NewReader(TEST_ID)

	read, err := readTrackId(reader, track)
	if err != nil {
		t.Errorf("Error reading track name %v", err)
	}

	if track.Id != expected {
		t.Errorf("Mismatched track Id.  Expected: %d.  Actual: %d", expected, track.Id)
	}

	if read != TRACK_ID_SIZE {
		t.Errorf("Mismatched bytes read. Expected %d.  Actual: %d", TRACK_ID_SIZE, read)
	}
}