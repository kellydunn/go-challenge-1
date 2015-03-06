// Package drum is supposed to implement the decoding of .splice drum machine files.
// See golang-challenge.com/go-challenge1/ for more information
package drum

import (
	"code.google.com/p/portaudio-go/portaudio"
	"encoding/binary"
	"fmt"
	"io"
	_ "io/ioutil"
	"os"
        "errors"
)

func Init() {
	portaudio.Initialize()
}

func LoadSample(filename string) (*Sample, error) {
	fd, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	sample := &Sample{
		Buffer: []float32{},
	}

	id, data, err := readChunk(fd)
	if err != nil {
		return nil, err
	}
	if id.String() != "FORM" {
		return nil, errors.New("Bad File Format")
	}
	_, err = data.Read(id[:])
	if err != nil {
		return nil, err
	}

	if id.String() != "AIFF" {
		return nil, errors.New("Bad File Format")
	}
	var c commonChunk
	var audio io.Reader
	for {
		id, chunk, err := readChunk(data)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		switch id.String() {
		case "COMM":
			err = binary.Read(chunk, binary.BigEndian, &c)
                        if err != nil {
                           return nil, err
                        }
                case "SSND":
			chunk.Seek(8, 1) //ignore offset and block
			audio = chunk
		default:
			fmt.Printf("ignoring unknown chunk '%s'\n", id)
		}
	}

        

	h, err := portaudio.DefaultHostApi()
	if err != nil {
		return nil, err
	}

	p := portaudio.LowLatencyParameters(h.DefaultInputDevice, h.DefaultOutputDevice)

	p.Input.Channels = 1
	p.Output.Channels = 1

	sample.Stream, err = portaudio.OpenStream(p, sample.Play)
	if err != nil {
		return nil, err
	}

	return sample, nil
}

type Sample struct {
	Stream *portaudio.Stream
	Buffer []float32
}

func (s *Sample) Play(in []float32, out []float32) {
	for i := range out {
		out[i] = s.Buffer[i]
	}

	fmt.Printf("FILLED BUFFER")
}

func readChunk(r readerAtSeeker) (id ID, data *io.SectionReader, err error) {
	_, err = r.Read(id[:])
	if err != nil {
		return
	}
	var n int32
	err = binary.Read(r, binary.BigEndian, &n)
	if err != nil {
		return
	}
	off, _ := r.Seek(0, 1)
	data = io.NewSectionReader(r, off, int64(n))
	_, err = r.Seek(int64(n), 1)
	return
}

type readerAtSeeker interface {
	io.Reader
	io.ReaderAt
	io.Seeker
}

type ID [4]byte

func (id ID) String() string {
	return string(id[:])
}

type commonChunk struct {
	NumChans      int16
	NumSamples    int32
	BitsPerSample int16
	SampleRate    [10]byte
}
