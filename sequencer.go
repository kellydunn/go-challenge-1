package drum

import ( "fmt" )

type Sequencer struct {
     Timer *Timer
     Bar int
     Beat int
}

func NewSequencer() *Sequencer {
     s := &Sequencer{
       Timer: NewTimer(120.0),
     }

     go func() {
        ppqnCount := 0
        for {
            select {
            case <- s.Timer.Pulses:
                 ppqnCount += 1
                 if ppqnCount % int(PPQN) == 0 {
                    fmt.Printf("Beat")
                 }

                 if ppqnCount == 15 {
                    ppqnCount = 0
                 }
            }           
        }
     }()

     s.Timer.Start()
     return s
}