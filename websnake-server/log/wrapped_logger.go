package log

import "log"

type wrappedLogger struct {
	logger  *log.Logger
	enabled bool
}

func (l wrappedLogger) Print(args ...interface{}) {
	if l.enabled {
		l.logger.Print(args...)
	}
}

func (l wrappedLogger) Printf(format string, args ...interface{}) {
	if l.enabled {
		l.logger.Printf(format, args...)
	}
}

func (l wrappedLogger) GetEnabled() bool {
	return l.enabled
}

func (l *wrappedLogger) SetEnabled(enabled bool) {
	l.enabled = enabled
}
