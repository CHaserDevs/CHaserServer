package utils

import(
  "io"
  "log"
)

type CustomLogger struct {
  infoLog     *log.Logger
  warningLog  *log.Logger
  importantLog *log.Logger
  errorLog    *log.Logger
}

func NewCustomLogger(output io.Writer) *CustomLogger {
  return &CustomLogger{
    infoLog:     log.New(output, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
    warningLog:  log.New(output, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
    importantLog: log.New(output, "IMPORTANT: ", log.Ldate|log.Ltime|log.Lshortfile),
    errorLog:    log.New(output, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
  }
}

func (l *CustomLogger) Info(message string) {
  l.infoLog.Println(message)
}

func (l *CustomLogger) Warning(message string) {
  l.warningLog.Println(message)
}

func (l *CustomLogger) Important(message string) {
  l.importantLog.Println(message)
}

func (l *CustomLogger) Error(err error) {
  l.errorLog.Println(err)
}
