package log

// NullLogger is a no-op instance of the Logger interface.
var NullLogger = nullLogger{}

// nullLogger implements a no-op type of the Logger interface.
type nullLogger struct{}

func (_ nullLogger) Emergency(args ...interface{})                 {}
func (_ nullLogger) Emergencyf(format string, args ...interface{}) {}
func (_ nullLogger) Alert(args ...interface{})                     {}
func (_ nullLogger) Alertf(format string, args ...interface{})     {}
func (_ nullLogger) Critical(args ...interface{})                  {}
func (_ nullLogger) Criticalf(format string, args ...interface{})  {}
func (_ nullLogger) Error(args ...interface{})                     {}
func (_ nullLogger) Errorf(format string, args ...interface{})     {}
func (_ nullLogger) Warning(args ...interface{})                   {}
func (_ nullLogger) Warningf(format string, args ...interface{})   {}
func (_ nullLogger) Notice(args ...interface{})                    {}
func (_ nullLogger) Noticef(format string, args ...interface{})    {}
func (_ nullLogger) Info(args ...interface{})                      {}
func (_ nullLogger) Infof(format string, args ...interface{})      {}
func (_ nullLogger) Debug(args ...interface{})                     {}
func (_ nullLogger) Debugf(format string, args ...interface{})     {}

func (_ nullLogger) Log(level Level, args ...interface{})                 {}
func (_ nullLogger) Logf(level Level, format string, args ...interface{}) {}

func (_ nullLogger) LogEmergency() bool        { return false }
func (_ nullLogger) LogAlert() bool            { return false }
func (_ nullLogger) LogCritical() bool         { return false }
func (_ nullLogger) LogError() bool            { return false }
func (_ nullLogger) LogWarning() bool          { return false }
func (_ nullLogger) LogNotice() bool           { return false }
func (_ nullLogger) LogInfo() bool             { return false }
func (_ nullLogger) LogDebug() bool            { return false }
func (_ nullLogger) LogLevel(level Level) bool { return false }

func (_ nullLogger) Close() error { return nil }
