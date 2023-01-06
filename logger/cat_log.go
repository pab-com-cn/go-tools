package logger

import (
  "io"
  "fmt"
)

type CatLog struct {
  outLog func(w io.Writer, file, msg string)
}

type logOption struct {
  level int
}

func NewCatLog() Log {
  cl := CatLog{} 
  return &cl
}

func printLog(level int, msg string) {
  var l string
  switch level {
case DebugLevel:
  l = "[DEBUG]"
case TraceLevel:
  l = "[TRACE]"
case InfoLevel:
  l = "[INFO]"
case WarnLevel:
  l = "[WARN]"
case ErrorLevel:
  l = "[ERROR]"
default:
  l = "[INFO]"
}
  fmt.Printf("%s", l)
}

func (cl *CatLog) IDebug(msg string) {}

func (cl *CatLog) IDebugF(format string, args ...interface{})

func (cl *CatLog) ITrace(msg string) {}

func (cl *CatLog) ITraceF(format string, args ...interface{})

func (cl *CatLog) IInfo(msg string) {}

func (cl *CatLog) IInfoF(format string, args ...interface{})

func (cl *CatLog) IWarn(msg string) {}

func (cl *CatLog) IWarnF(format string, args ...interface{})

func (cl *CatLog) IError(msg string) {}

func (cl *CatLog) IErrorF(format string, args ...interface{})

func (cl *CatLog) IFatal(err error) {}

func (cl *CatLog) IFatalF(format string, args ...interface{})
