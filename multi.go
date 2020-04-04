package mixers

import (
	logging "github.com/codemodify/systemkit-logging"
)

type multiLogger struct {
	loggers []logging.Logger
}

// NewMulti -
func NewMulti(loggers []logging.Logger) logging.Logger {
	return &multiLogger{
		loggers: loggers,
	}
}

func (thisRef multiLogger) Log(logEntry logging.LogEntry) logging.LogEntry {
	for _, logger := range thisRef.loggers {
		logger.Log(logEntry)
	}

	return logging.LogEntry{}
}
