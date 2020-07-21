package logz

type prefixedLogger struct {
	main   Logger
	prefix string
}

func (l *prefixedLogger) Debugf(format string, args ...interface{}) {
	l.main.Debugf(l.prefix+format, args...)
}

func (l *prefixedLogger) Errorf(format string, args ...interface{}) {
	l.main.Errorf(l.prefix+format, args...)
}

func (l *prefixedLogger) Infof(format string, args ...interface{}) {
	l.main.Infof(l.prefix+format, args...)
}

func (l *prefixedLogger) Warningf(format string, args ...interface{}) {
	l.main.Warningf(l.prefix+format, args...)
}

func NewPrefixedLogger(main Logger, prefix string) Logger {
	return &prefixedLogger{
		main:   main,
		prefix: prefix,
	}
}
