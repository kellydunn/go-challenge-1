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

type Timer struct {
     Pulses chan int
     Done chan bool
     Lock *sync.Mutex
     BPM float64
}

func NewTimer(bpm float64) *Timer {
     t := &Timer{
       Pulses: make(chan int),
       Done: make(chan bool),
       BPM: bpm,
     }

     return t
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
              fmt.Printf("PULSE\n")
            }
        }
     }()
}

func microsecondsPerPulse(bpm float64) time.Duration {
	return time.Duration((MINUTE * MICROSECOND) / (PPQN * bpm))
}
