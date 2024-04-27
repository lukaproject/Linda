package task

import "os"

func SignalMsg(s os.Signal) string {
	return "signal: " + s.String()
}
