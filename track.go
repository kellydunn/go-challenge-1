package drum

import (
	"encoding/binary"
	"io"
)

type Track struct {
	ID           uint8
	NameLen      uint32
	Name         string
	Sample       *Sample
	Buffer       []float32
	Playhead     int
	StepSequence StepSequence
}

func readTrackID(reader io.Reader, t *Track) (int, error) {
	var id uint8
	err := binary.Read(reader, binary.BigEndian, &id)
	if err != nil {
		return 0, err
	}

	t.ID = id

	return TrackIDSize, nil
}

func readTrackStepSequence(reader io.Reader, t *Track) (int, error) {
	steps := make([]byte, StepSequenceSize)
	err := binary.Read(reader, binary.BigEndian, &steps)
	if err != nil {
		return 0, err
	}

	t.StepSequence = StepSequence{Steps: steps}

	return StepSequenceSize, nil
}

func readTrackName(reader io.Reader, t *Track) (int, error) {
	bytesRead := 0

	var trackNameLen uint32
	err := binary.Read(reader, binary.BigEndian, &trackNameLen)
	if err != nil {
		return bytesRead, err
	}

	bytesRead += TrackNameSize

	trackNameBytes := make([]byte, trackNameLen)
	err = binary.Read(reader, binary.BigEndian, trackNameBytes)
	if err != nil {
		return bytesRead, err
	}

	t.Name = string(trackNameBytes)

	bytesRead += int(trackNameLen)

	return bytesRead, nil
}
