package xlog

import (
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Debug(args ...any)
	Debugf(format string, args ...any)

	Info(args ...any)
	Infof(format string, args ...any)

	Warn(args ...any)
	Warnf(format string, args ...any)

	Error(args ...any)
	Errorf(format string, args ...any)

	Fatal(args ...any)
	Fatalf(format string, args ...any)
}

func Initial() {
	setUpForDev()
}

func NewForPackage() Logger {
	pc, _, _, _ := runtime.Caller(1)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	packageName := ""
	if parts[pl-2][0] == '(' {
		packageName = strings.Join(parts[0:pl-2], ".")
	} else {
		packageName = strings.Join(parts[0:pl-1], ".")
	}
	return logrus.WithField("package", packageName)
}

func setUpForDev() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}
