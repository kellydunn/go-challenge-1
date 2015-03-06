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
	SAMPLE_RATE       = 44100
	FRAMES_PER_SAMPLE = 256
	INPUT_CHANNELS    = 0
	OUTPUT_CHANNELS   = 2
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
	buffer := make([]float32, FRAMES_PER_BUFFER)
	s := &Sample{
		Buffer: buffer,
	}

	s.Stream, err = portaudio.OpenDefaultStream(
		INPUT_CHANNELS,
		OUTPUT_CHANNELS,
		SAMPLE_RATE,
		FRAMES_PER_SAMPLE,
		&s.Buffer,
	)

        if err != nil {
		return nil, err
	}

	s.Stream.Start()
	defer s.Stream.Close()

	_, err = soundFile.ReadItems(buffer)
	if err != nil {
		return nil, err
	}

	s.Stream.Write()

	time.Sleep(time.Second * 5)

	s.Stream.Stop()

	return s, nil
}

type Sample struct {
	Buffer []float32
	Stream *portaudio.Stream
}
