package xlog

import "github.com/sirupsen/logrus"

func Debug(args ...any)                 { logrus.Debug(args...) }
func Debugf(format string, args ...any) { logrus.Debugf(format, args...) }
func Info(args ...any)                  { logrus.Info(args...) }
func Infof(format string, args ...any)  { logrus.Infof(format, args...) }
func Warn(args ...any)                  { logrus.Warn(args...) }
func Warnf(format string, args ...any)  { logrus.Warnf(format, args...) }
func Error(args ...any)                 { logrus.Error(args...) }
func Errorf(format string, args ...any) { logrus.Errorf(format, args...) }
func Fatal(args ...any)                 { logrus.Fatal(args...) }
func Fatalf(format string, args ...any) { logrus.Fatalf(format, args...) }
