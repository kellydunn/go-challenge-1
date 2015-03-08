package drum

import (
	"encoding/binary"
	"bytes"
	"io"
	"fmt"
)

// Pattern is the high level representation of the
// drum pattern contained in a .splice file.
type Pattern struct {
	Version string
	Tempo   float32
	Tracks  []*Track
}

func NewPattern(reader io.Reader) (*Pattern, error) {
	p := &Pattern{}

	// Header
	spliceFile := make([]byte, SPLICE_FILE_SIZE)
	err := binary.Read(reader, binary.BigEndian, spliceFile)
	if err != nil {
		return nil, err
	}

	var size uint64
	err = binary.Read(reader, binary.BigEndian, &size)
	if err != nil {
		return nil, err
	}

	// Keeps track of all the bytes read after the SPLICE string
	// So it will ignore garbage data after the relevant information.
	bytesRead := 0

	// Read Version
	read, err := readPatternVersionString(reader, p)
	if err != nil {
		return nil, err
	}
	
	bytesRead += read

	// Read Tempo
	read, err = readPatternTempo(reader, p)
	if err != nil {
		return nil, err
	}
	
	bytesRead += read

	p.Tracks = []*Track{}

	for bytesRead < int(size) {
		// Read tracks
		t := &Track{}

		read, err = readTrackId(reader, t)
		if err != nil {
			return nil, err
		}
		
		bytesRead += read

		read, err = readTrackName(reader, t)
		if err != nil {
			return nil, err
		}
		
		bytesRead += read
		
		read, err = readTrackStepSequence(reader, t)
		if err != nil {
			return  nil, err
		}
		
		bytesRead += read

		p.Tracks = append(p.Tracks, t)
	}

	return p, nil
}

func (p *Pattern) String() string {
	buf := bytes.NewBufferString("")
	buf.WriteString(fmt.Sprintf("Saved with HW Version: %s\n", p.Version))
	buf.WriteString(fmt.Sprintf("Tempo: %3v\n", p.Tempo))

	for _, track := range p.Tracks {
		buf.WriteString(fmt.Sprintf("(%d) %s\t%s\n", track.Id, track.Name, track.StepSequence))
	}

	return buf.String()
}

func readPatternVersionString(reader io.Reader, p *Pattern) (int, error) {
	version := make([]byte, VERSION_SIZE)
	err := binary.Read(reader, binary.BigEndian, version)
	if err != nil {
		return 0, err
	}

	p.Version = string(bytes.Trim(version, EMPTY_BYTE))
	
	return VERSION_SIZE, nil
}

func readPatternTempo(reader io.Reader, p*Pattern) (int, error) {
	var tempo float32
	err := binary.Read(reader, binary.LittleEndian, &tempo)
	if err != nil {
		return 0, err
	}

	p.Tempo = tempo

	return BPM_SIZE, nil
}