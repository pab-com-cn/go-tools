package logger

const (
  DebugLevel = iota
  TraceLevel
  InfoLevel
  WarnLevel
  ErrorLevel
)

type Log interface {
  IDebug(msg string)
  IDebugF(format string, args ...interface{})
  ITrace(msg string)
  ITraceF(format string, args ...interface{})
  IInfo(msg string)
  IInfoF(format string, args ...interface{})
  IWarn(msg string)
  IWarnF(format string, args ...interface{})
  IError(msg string)
  IErrorF(format string, args ...interface{})
  IFatal(err error)
  IFatalF(format string, args ...interface{})
}
