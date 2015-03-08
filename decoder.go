package drum

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var SPLICE_FILE_SIZE int = 6
var FILE_SIZE int = 8
var VERSION_SIZE int = 32
var BPM_SIZE int = 4
var TRACK_ID_SIZE int = 1
var TRACK_NAME_SIZE int = 4
var STEP_SEQUENCE_SIZE = 16
var EMPTY_BYTE = "\x00"

// DecodeFile decodes the drum machine file found at the provided path
// and returns a pointer to a parsed pattern which is the entry point to the
// rest of the data.
func DecodeFile(path string) (*Pattern, error) {
	
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	p := &Pattern{}

	reader := bytes.NewReader(data)

	spliceFile := make([]byte, SPLICE_FILE_SIZE)
	err = binary.Read(reader, binary.BigEndian, spliceFile)
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

	version := make([]byte, VERSION_SIZE)
	err = binary.Read(reader, binary.BigEndian, version)
	if err != nil {
		return nil, err
	}
	bytesRead += VERSION_SIZE

	var bpm float32
	err = binary.Read(reader, binary.LittleEndian, &bpm)
	if err != nil {
		return nil, err
	}
	bytesRead += BPM_SIZE

	p.Version = string(bytes.Trim(version, EMPTY_BYTE))
	p.Tempo = bpm
	p.Tracks = []Track{}

	for bytesRead < int(size) {
		// Read tracks
		t := Track{}

		var id uint8
		err = binary.Read(reader, binary.BigEndian, &id)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		bytesRead += 1

		var trackNameLen uint32
		err = binary.Read(reader, binary.BigEndian, &trackNameLen)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		bytesRead += 4

		trackNameBytes := make([]byte, trackNameLen)
		err = binary.Read(reader, binary.BigEndian, trackNameBytes)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		trackName := string(trackNameBytes)
		bytesRead += int(trackNameLen)

		steps := make([]byte, STEP_SEQUENCE_SIZE)
		err = binary.Read(reader, binary.BigEndian, &steps)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		bytesRead += STEP_SEQUENCE_SIZE

		t.Id = id
		t.Name = trackName
		t.StepSequence = StepSequence{Steps: steps}
		t.Sample, err = LoadSample("kits/" + p.Version + "/" + t.Name + ".wav")

		if err != nil {
			return nil, err
		}

		p.Tracks = append(p.Tracks, t)
	}

	return p, nil
}

// Pattern is the high level representation of the
// drum pattern contained in a .splice file.
// TODO: implement
type Pattern struct {
	Version string
	Tempo   float32
	Tracks  []Track
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

type Track struct {
	Id           uint8
	NameLen      uint32
	Name         string
	Sample       *Sample
	StepSequence StepSequence
}

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
