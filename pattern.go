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

// Creates and returns a pointer to a new Pattern.
// Returns an error if one was encountered from attempting to read
// from the passed in io.Reader.  This can happen if the
// passed in reader does not comply with the binary Splice format.
//
// To learn more about the format, visit the README.md at the root level of this project.
func NewPattern(reader io.Reader) (*Pattern, error) {
	p := &Pattern{}

	// Read the "SPLICE" header from the binary file.
	spliceFile := make([]byte, SPLICE_FILE_SIZE)
	err := binary.Read(reader, binary.BigEndian, spliceFile)
	if err != nil {
		return nil, err
	}

	// Read the size of the file from the header.
	var size uint64
	err = binary.Read(reader, binary.BigEndian, &size)
	if err != nil {
		return nil, err
	}

	// This counter keeps track of all the bytes read after the headers
	// such that it will ignore garbage data that occurs after the designated size.
	bytesRead := 0

	// Reads the Pattern Version from the passed in binary file.
	read, err := readPatternVersionString(reader, p)
	if err != nil {
		return nil, err
	}
	
	bytesRead += read

	// Reads the Pattern Tempo from the passed in binary file.	
	read, err = readPatternTempo(reader, p)
	if err != nil {
		return nil, err
	}
	
	bytesRead += read

	p.Tracks = []*Track{}

	// Until we read up to the passed in size of the file
	// We will continue to consume data as Tracks.
	for bytesRead < int(size) {

		t := &Track{}

		// Reads the Track Id from the passed in binary file.			
		read, err = readTrackId(reader, t)
		if err != nil {
			return nil, err
		}
		
		bytesRead += read

		// Reads the Track Name from the passed in binary file.		
		read, err = readTrackName(reader, t)
		if err != nil {
			return nil, err
		}
		
		bytesRead += read
		
		// Reads the Track Step sequence from the passed in binary file.
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