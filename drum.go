// Package drum is supposed to implement the decoding of .splice drum machine files.
// See golang-challenge.com/go-challenge1/ for more information
package drum

import (
	"code.google.com/p/portaudio-go/portaudio"
	"github.com/mkb218/gosndfile/sndfile"
        "time"
)

const (
	FRAMES_PER_BUFFER = 8196
)

func Init() {
	portaudio.Initialize()
}

func LoadSample(filename string) (*Sample, error) {
	var info sndfile.Info
	soundFile, err := sndfile.Open(filename, sndfile.Read, &info)
	if err != nil {
		return nil, err
	}

	defer soundFile.Close()
	buffer := make([]float32, 10*info.Samplerate*info.Channels)        
	s := &Sample{
		Buffer: buffer,
	}

	s.Stream, err = portaudio.OpenDefaultStream(0, 1, 44100, len(s.Buffer), &s.Buffer)
	if err != nil {
		return nil, err
	}

        s.Stream.Start()
        defer s.Stream.Close()

	_, err = soundFile.ReadItems(buffer)
	if err != nil {
		return nil, err
	}

        time.Sleep(time.Second * 10)

        return s, nil
}

func (s *Sample) Play(in []float32, out []float32) {
	for i := range in {
		out[i] = in[i]
	}
}

type Sample struct {
	Buffer    []float32
	Stream *portaudio.Stream
}
