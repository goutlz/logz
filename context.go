package logz

import "context"

const (
	ctxLoggerKey = "ctx-logger"
)

var (
	defaultLogger Logger
)

func SetLogger(ctx context.Context, logger Logger) context.Context {
	if ctx == nil {
		return context.WithValue(context.Background(), ctxLoggerKey, logger)
	}
	return context.WithValue(ctx, ctxLoggerKey, logger)
}

func GetLogger(ctx context.Context) Logger {
	if ctx == nil {
		return defaultLogger
	}
	v := ctx.Value(ctxLoggerKey)
	if v == nil {
		return defaultLogger
	}
	res, ok := v.(Logger)
	if !ok {
		return defaultLogger
	}
	return res
}

func init() {
	defaultLogger = NewStdLogger()
}
