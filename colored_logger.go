package logz

import "log"

const (
	resetColorCode = "\033[0m"

	debugColorCode   = "\033[38;5;2m"
	errorColorCode   = "\033[38;5;1m"
	warningColorCode = "\033[38;5;3m"
	infoColorCode    = "\033[38;5;4m"
)

type coloredLogger struct{}

func (l *coloredLogger) Debugf(format string, args ...interface{}) {
	log.Printf(debugColorCode+debug_prefix+resetColorCode+format, args...)
}

func (l *coloredLogger) Errorf(format string, args ...interface{}) {
	log.Printf(errorColorCode+error_prefix+resetColorCode+format, args...)
}

func (l *coloredLogger) Infof(format string, args ...interface{}) {
	log.Printf(infoColorCode+info_prefix+resetColorCode+format, args...)
}

func (l *coloredLogger) Warningf(format string, args ...interface{}) {
	log.Printf(warningColorCode+warning_prefix+resetColorCode+format, args...)
}

func NewColoredLogger() Logger {
	return &coloredLogger{}
}
