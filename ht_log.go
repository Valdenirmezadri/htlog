package htl

import (
	"os"

	logging "github.com/op/go-logging"
	"gopkg.in/natefinch/lumberjack.v2"
)

var to *log

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

	to = &log{
		Logger:  logging.MustGetLogger("ht-log"),
		Options: o,
	}

	if to.mode == "dev" {
		to.close, err = to.devLog()
		if err != nil {
			return err
		}

		return nil
	}

	to.close, err = to.prodLog()

	return err
}

func Stop() error {
	if to.close == nil {
		return nil
	}

	return to.close()
}

func Log() *logging.Logger {
	return to.Logger
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
