package drum

import (
	"code.google.com/p/portaudio-go/portaudio"
_	"fmt"
)

type Sequencer struct {
	Timer    *Timer
	Bar      int
	Beat     int
	Pattern  *Pattern
	Stream   *portaudio.Stream
}

func NewSequencer() *Sequencer {
	s := &Sequencer{
		Timer:    NewTimer(),
		Bar:      0,
		Beat:     0,
	}
	
	stream, err := portaudio.OpenDefaultStream(
		0,
		2,
		44100,
		portaudio.FramesPerBufferUnspecified,
		s.ProcessAudio,
		)

	if err != nil {
		panic(err)
	}

	s.Stream = stream

	return s
}

func (s *Sequencer) Start() {
	go func() {
		ppqnCount := 0

		for {
			select {
			case <-s.Timer.Pulses:
				ppqnCount += 1

				// TODO add in time signatures
				if ppqnCount%(int(PPQN)/4) == 0 {
					go s.PlayTrigger()

					s.Beat += 1
					s.Beat = s.Beat % 4
				}

				// TODO Add in time signatures
				if ppqnCount%int(PPQN) == 0 {
					s.Bar += 1
					s.Bar = s.Bar % 4
				}

				// 4 bars of quarter notes
				if ppqnCount == (int(PPQN) * 4) {
					ppqnCount = 0
				}

			}
		}
	}()

	s.Timer.Start()
	s.Stream.Start()
}

func (s *Sequencer) ProcessAudio(out []float32) {
	for i := range out {
		var data float32
		
		for _, track := range s.Pattern.Tracks {
			if track.Sample.Playhead < len(track.Sample.Buffer) {
				data += track.Sample.Buffer[track.Sample.Playhead]
				track.Sample.Playhead++
			}
		}

		if data > 1.0 {
			data = 1.0
		}

		out[i] = data
	}
}

func (s * Sequencer) PlayTrigger() {
	index := (s.Bar * 4) + s.Beat
	
	for _, track := range s.Pattern.Tracks {
		if track.StepSequence.Steps[index] == byte(1) {
			track.Sample.Playhead = 0
		}
	}
}
	