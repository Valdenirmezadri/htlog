package htl

import (
	"os"

	logging "github.com/op/go-logging"
	"gopkg.in/natefinch/lumberjack.v2"
)

var _log *log

type log struct {
	Options
	*logging.Logger
	close func() error
}

func Start(ops ...Optfunc) (err error) {
	o := defaultOps()
	for _, fn := range ops {
		fn(&o)
	}

	_log = &log{
		Logger:  logging.MustGetLogger("ht-log"),
		Options: o,
	}

	if _log.mode == "dev" {
		_log.close, err = _log.devLog()
		if err != nil {
			return err
		}

		return nil
	}

	_log.close, err = _log.prodLog()

	return err
}

func Stop() error {
	if _log.close == nil {
		return nil
	}

	return _log.close()
}

func Default() *log {
	return _log
}

func (l *log) devLog() (close func() error, err error) {
	console := logging.NewLogBackend(os.Stderr, "", 0)
	consoleFormatter := logging.NewBackendFormatter(console, formatConsole)

	writer, close, err := l.writerToWithRotation()
	if err != nil {
		return nil, err
	}

	fileFormatter := logging.NewBackendFormatter(writer, formatFile)
	logging.SetBackend(consoleFormatter, fileFormatter)
	return close, nil
}

func (l *log) prodLog() (close func() error, err error) {
	writer, close, err := l.writerToWithRotation()
	if err != nil {
		return nil, err
	}

	fileFormatter := logging.NewBackendFormatter(writer, formatFile)
	logging.SetBackend(fileFormatter)
	return close, nil
}

func (l *log) writerToWithRotation() (writer *logging.LogBackend, close func() error, err error) {
	rotate := &lumberjack.Logger{
		Filename:   l.pathLog,
		MaxSize:    int(l.maxAge),
		MaxBackups: int(l.maxBackups),
		MaxAge:     int(l.maxAge),
		Compress:   l.compress,
	}

	return logging.NewLogBackend(rotate, "", 0), rotate.Close, nil
}

var formatConsole = logging.MustStringFormatter(
	`%{color} %{time:15:04:05.000} %{shortfile} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

var formatFile = logging.MustStringFormatter(
	`%{time:Jan 02 2006 15:04:05} %{shortfile} ▶ %{level:.4s} %{message}`,
)
