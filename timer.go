package drum

import "time"

const (
	PPQN        = 24.0
	MINUTE      = 60.0
	MICROSECOND = 1000000000
)

func microsecondsPerPulse(bpm float64) time.Duration {
	return time.Duration((MINUTE * MICROSECOND) / (PPQN * bpm))
}
