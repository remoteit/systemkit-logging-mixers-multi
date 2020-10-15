package mixers

import (
	logging "github.com/remoteit/systemkit-logging"
)

type multiLogger struct {
	loggers []logging.CoreLogger
}

// NewMultiLogger -
func NewMultiLogger(loggers []logging.CoreLogger) logging.CoreLogger {
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
