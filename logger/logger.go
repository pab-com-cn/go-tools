package logger

var log Log

func init() {
  log = logger.NewCatLog()
}

func Debug(msg string) {}

func DebugF(format string, args ...interface{}) {}

func Tarce(msg string) {}

func TarceF(format string, args ...interface{}) {}

func Info(msg string) {}

func InfoF(format string, args ...interface{}) {}

func Warn(msg string) {}

func WarnF(format string, args ...interface{}) {}

func Error(msg string) {}

func ErrorF(format string, args ...interface{}) {}

func Fatal(err error) {}

func FatalF(format string, args ...interface{}) {}
