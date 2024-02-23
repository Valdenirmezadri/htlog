package htl

// args ...interface{}
// args...

// format string, args ...interface{}
// format, args...

// Info logs a message using INFO as log level.
func Info(args ...interface{}) {
	_log.Info(args...)
}

// Debug logs a message using DEBUG as log level.
func Debug(args ...interface{}) {
	_log.Debug(args...)
}

// Warning logs a message using WARNING as log level.
func Warning(args ...interface{}) {
	_log.Warning(args...)
}

func Error(args ...interface{}) {
	_log.Error(args...)
}

func Critical(args ...interface{}) {
	_log.Critical(args...)
}

func Panic(args ...interface{}) {
	_log.Panic(args...)
}

func Fatal(args ...interface{}) {
	_log.Fatal(args...)
}

func Infof(format string, args ...interface{}) {
	_log.Infof(format, args...)
}

func Debugf(format string, args ...interface{}) {
	_log.Debugf(format, args...)
}

func Warningf(format string, args ...interface{}) {
	_log.Warningf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	_log.Errorf(format, args...)
}

func Criticalf(format string, args ...interface{}) {
	_log.Criticalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	_log.Panicf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	_log.Fatalf(format, args...)
}
