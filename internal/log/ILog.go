package log

type ILog interface {
	Debug(string, ...interface{})
	Info(string, ...interface{})
	Warn(string, ...interface{})
	Error(error, ...interface{})
	Panic(string, ...interface{})
	Fatal(string, ...interface{})
	//需要格式化日志 ，最后一个是context
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(error, ...interface{})
	Panicf(string, ...interface{})
	Fatalf(string, ...interface{})
}
