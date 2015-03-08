package drum

//import ( "fmt" )

type Sequencer struct {
	Timer *Timer
	Bar int
	Beat int
	Pattern *Pattern
}

func NewSequencer() *Sequencer {
	s := &Sequencer{
		Timer: NewTimer(),
		Bar: 0,
		Beat: 0,
	}
	
	return s
}

func (s * Sequencer) Start() {
	go func() {
		ppqnCount := 0
		
		for {
			select {
			case <- s.Timer.Pulses:
				ppqnCount += 1

				// TODO add in time signatures
				if ppqnCount % (int(PPQN)/4) == 0 {
					s.PlayTriggers()
					
					s.Beat += 1
					s.Beat = s.Beat % 4
				}

				// TODO Add in time signatures
				if ppqnCount % int(PPQN) == 0 {
					s.Bar += 1
					s.Bar = s.Bar % 4
				}

				// 4 bars of quarter notes
				if ppqnCount == (int(PPQN) * 4)  {
					ppqnCount = 0
				}				
				
			}
		}
	}()
		
	s.Timer.Start()
}

func (s * Sequencer) PlayTriggers() {
	for _, track := range s.Pattern.Tracks {
		// TODO time signatures
		
		//fmt.Printf("Checking index: %d for track: %s.\tSequence: %s\n", (s.Bar * 4) + s.Beat, track.Name, track.StepSequence)
		if track.StepSequence.Steps[(s.Bar * 4) + s.Beat] == byte(1) {
			track.Sample.Play()
		}
	}
}