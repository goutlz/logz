package logz

import "log"

type stdLogger struct{}

func (l *stdLogger) Debugf(format string, args ...interface{}) {
	log.Printf(debug_prefix+format, args...)
}

func (l *stdLogger) Errorf(format string, args ...interface{}) {
	log.Printf(error_prefix+format, args...)
}

func (l *stdLogger) Infof(format string, args ...interface{}) {
	log.Printf(info_prefix+format, args...)
}

func (l *stdLogger) Warningf(format string, args ...interface{}) {
	log.Printf(warning_prefix+format, args...)
}

func NewStdLogger() Logger {
	return &stdLogger{}
}
