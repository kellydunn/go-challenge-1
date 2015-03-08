// Package drum is supposed to implement the decoding of .splice drum machine files.
// See golang-challenge.com/go-challenge1/ for more information
package drum

import (
	"code.google.com/p/portaudio-go/portaudio"
	"github.com/mkb218/gosndfile/sndfile"
)

const (
	FRAMES_PER_BUFFER = 12800
	SAMPLE_RATE       = 44100
	FRAMES_PER_SAMPLE = 256
	INPUT_CHANNELS    = 0
	OUTPUT_CHANNELS   = 2
)

var Master *Sequencer

func Init() {
	portaudio.Initialize()
	Master = NewSequencer()
}

func LoadSample(filename string) (*Sample, error) {
	var info sndfile.Info
	soundFile, err := sndfile.Open(filename, sndfile.Read, &info)
	if err != nil {
		return nil, err
	}

	buffer := make([]float32, 10*info.Samplerate*info.Channels)
	numRead, err := soundFile.ReadItems(buffer)
	if err != nil {
		return nil, err
	}

	defer soundFile.Close()

	s := &Sample{
		Buffer: buffer[:numRead],
		Playhead: int(numRead),
	}

	if err != nil {
		return nil, err
	}

	return s, nil
}

type Sample struct {
	Buffer []float32
	Playhead int
}
