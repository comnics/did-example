package util

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Logger
type Logger struct {
	Trace *log.Logger
	Warn  *log.Logger
	Info  *log.Logger
	Error *log.Logger
}

func GetLogger() Logger {
	return InitLog(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
}

// InitLog: Initialize the log
func InitLog(traceHandle io.Writer, infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer) Logger {
	var logger Logger

	logger.Trace = log.New(traceHandle, "[TRACE] ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Info = log.New(infoHandle, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Warn = log.New(warningHandle, "[WARNING] ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Error = log.New(errorHandle, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)

	return logger
}
