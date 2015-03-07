package drum

import ("time"
       "sync"
       "fmt"
       "log")

const (
	PPQN        = 24.0
	MINUTE      = 60.0
	MICROSECOND = 1000000000
)

var DEFAULT_BPM float32 = 120.0

type Timer struct {
     Pulses chan int
     Done chan bool
     Lock *sync.Mutex
     BPM float32
}

func NewTimer() *Timer {
     t := &Timer{
       Pulses: make(chan int),
       Done: make(chan bool),
       BPM: DEFAULT_BPM,
     }

     return t
}

func (t * Timer) SetBPM(bpm float32) {
	t.BPM = bpm
}

func (t * Timer) Start() {
     log.Printf("STARTING TIMER")
     go func() {
        for {
            select {
            case <- t.Done :
              fmt.Printf("DONE\n")
              break
            default:
              interval := microsecondsPerPulse(t.BPM)
              time.Sleep(interval)
              t.Pulses <- 1
            }
        }
     }()
}

func microsecondsPerPulse(bpm float32) time.Duration {
	return time.Duration((MINUTE * MICROSECOND) / (PPQN * bpm))
}
