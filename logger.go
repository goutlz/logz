package logz

const (
	debug_prefix   = "[DEBUG] "
	error_prefix   = "[ERROR] "
	info_prefix    = "[INFO] "
	warning_prefix = "[WARNING] "
)

type Logger interface {
	Debugf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warningf(format string, args ...interface{})
}
