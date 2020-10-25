/*******************************************************************************
实现：日志接口
作者：Lemine
时间：2020/09/09
*******************************************************************************/
package loggo

type Logger interface {
	Debug(msg interface{})
	FormatDebug(format string, args ...interface{})

	Trace(msg interface{})
	FormatTrace(format string, args ...interface{})

	Info(msg interface{})
	FormatInfo(format string, args ...interface{})

	Warn(msg interface{})
	FormatWarn(format string, args ...interface{})

	Error(msg interface{})
	FormatError(format string, args ...interface{})

	Fatal(msg interface{})
	FormatFatal(format string, args ...interface{})

	Close()
}
