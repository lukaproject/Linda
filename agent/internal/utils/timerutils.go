package utils

import "time"

const (
	defaultTickDuration = 500 * time.Millisecond
)

type TimerUtils struct{}

func (tu TimerUtils) WaitUntil(condFunc func() bool) {}

// true for success reached condFunc come true,
// false for timeout
func (tu TimerUtils) WaitUntilWithTimeout(condFunc func() bool, timeout time.Duration) bool {
	ticker := time.NewTicker(defaultTickDuration)
	timeoutPoint := time.Now().Add(timeout)
	for range ticker.C {
		if condFunc() {
			return true
		}
		if time.Now().After(timeoutPoint) {
			return false
		}
	}
	return false
}
