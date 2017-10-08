package clock

import (
	"time"
)

type Clock interface {
	Now() time.Time
	Ticker() <-chan time.Time
}
