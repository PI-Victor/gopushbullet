package log

import (
	"os"

	"github.com/op/go-logging"
)

const format = "%{color}%{time:15:04:05.000}%{level:.6s} ▶ %{color:reset} %{message}"

var (
	log           = logging.MustGetLogger("gopush")
	formatter     = logging.MustStringFormatter(format)
)

func init() {
	infoBackend := logging.NewLogBackend(os.Stderr, "", 0)
	errBackend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(infoBackend, formatter)

	// Send just errors to stderr
	errLeveled := logging.AddModuleLevel(errBackend)
	errLeveled.SetLevel(logging.ERROR, "gopush")
	logging.SetBackend(errLeveled, backendFormatter)
}

func Debug(format string, args ...interface{}){log.Debugf(format, args...)}
func Info(format string, args ...interface{}){log.Infof(format, args...)}
func Notice(format string, args ...interface{}){log.Noticef(format, args...)}
func Warning(format string, args ...interface{}){log.Warningf(format, args...)}
func Error(format string, args ...interface{}){log.Errorf(format, args...)}
func Critical(format string, args ...interface{}){log.Criticalf(format, args...)}
func Panic(format string, args ...interface{}){log.Panicf(format, args...)}

