package drum

import (
	"bytes"
	"testing"
)

var TEST_ID []byte = []byte{0x01}

var TEST_NAME []byte = []byte{
	0x00, 0x00, 0x00, 0x04, 0x6b, 0x69, 0x63, 0x6b,
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
		t.Errorf("Error reading track Id %v", err)
	}

	if track.Id != expected {
		t.Errorf("Mismatched track Id.  Expected: %d.  Actual: %d", expected, track.Id)
	}

	if read != TRACK_ID_SIZE {
		t.Errorf("Mismatched bytes read. Expected %d.  Actual: %d", TRACK_ID_SIZE, read)
	}
}

func TestReadTrackName(t *testing.T) {
	track := &Track{}
	expected := "kick"
	reader := bytes.NewReader(TEST_NAME)

	read, err := readTrackName(reader, track)
	if err != nil {
		t.Errorf("Error reading track name %v", err)
	}

	if track.Name != expected {
		t.Errorf("Mismatched track Name.  Expected: %s.  Actual: %s", expected, track.Name)
	}

	if read != TRACK_NAME_SIZE+len(expected) {
		t.Errorf("Mismatched bytes read. Expected %d.  Actual: %d", TRACK_NAME_SIZE+len(expected), read)
	}
}

func TestReadStepSequence(t *testing.T) {
	track := &Track{}
	expected := TEST_STEP_SEQUENCE
	reader := bytes.NewReader(TEST_STEP_SEQUENCE)

	read, err := readTrackStepSequence(reader, track)
	if err != nil {
		t.Errorf("Unable to reat step sequence")
	}

	for i := range expected {
		if track.StepSequence.Steps[i] != expected[i] {
			t.Errorf("Mismatched track StepSequence steps.  Expected: %v.  Actual: %v", expected[i], track.StepSequence.Steps[i])
		}
	}

	if read != STEP_SEQUENCE_SIZE {
		t.Errorf("Mismatched bytes read. Expected %d.  Actual: %d", STEP_SEQUENCE_SIZE, read)
	}

}
