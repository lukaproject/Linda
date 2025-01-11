package xos

import (
	"os"
	"os/signal"
)

func WaitForSignal(signals ...os.Signal) os.Signal {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, signals...)
	return <-ch
}
