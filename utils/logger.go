package utils

import (
	"fmt"
	"log"
	"os"
)

/* Summary
Here is Logger system for output

LogLevel:
  Error
  Warn
  Info
  Debug

It need log level in init this log system

Please import in all code and always output log with this.
*/

type Logger struct {
	*log.Logger
	level LogLevel
}

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

func NewLogger(level LogLevel) *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile),
		level:  level,
	}
}

func (l *Logger) log(level LogLevel, msg string) {
	if level >= l.level {
		var prefix string
		switch level {
		case DEBUG:
			prefix = "DEBUG"
		case INFO:
			prefix = "INFO"
		case WARN:
			prefix = "WARN"
		case ERROR:
			prefix = "ERROR"
		}
		l.Logger.SetPrefix(fmt.Sprintf("[%s] ", prefix))
		l.Logger.Println(msg)
	}
}

func (l *Logger) Debug(msg string) {
	l.log(DEBUG, msg)
}

func (l *Logger) Info(msg string) {
	l.log(INFO, msg)
}

func (l *Logger) Warn(msg string) {
	l.log(WARN, msg)
}

func (l *Logger) Error(msg string) {
	l.log(ERROR, msg)
}
