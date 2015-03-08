package drum

import (
	"encoding/binary"
	"io"
)

type Track struct {
	Id           uint8
	NameLen      uint32
	Name         string
	Sample       *Sample
	StepSequence StepSequence
}

func readTrackId(reader io.Reader, t *Track) (int, error) {
	var id uint8
	err := binary.Read(reader, binary.BigEndian, &id)
	if err != nil {
		return 0, err
	}

	t.Id = id

	return TRACK_ID_SIZE, nil
}

func readTrackStepSequence(reader io.Reader, t *Track) (int, error) {
	steps := make([]byte, STEP_SEQUENCE_SIZE)
	err := binary.Read(reader, binary.BigEndian, &steps)
	if err != nil {
		return 0, err
	}

	t.StepSequence = StepSequence{Steps: steps}

	return STEP_SEQUENCE_SIZE, nil
}

func readTrackName(reader io.Reader, t *Track) (int, error) {
	bytesRead := 0

	var trackNameLen uint32
	err := binary.Read(reader, binary.BigEndian, &trackNameLen)
	if err != nil {
		return bytesRead, err
	}

	bytesRead += TRACK_NAME_SIZE

	trackNameBytes := make([]byte, trackNameLen)
	err = binary.Read(reader, binary.BigEndian, trackNameBytes)
	if err != nil {
		return bytesRead, err
	}

	t.Name = string(trackNameBytes)

	bytesRead += int(trackNameLen)

	return bytesRead, nil
}
