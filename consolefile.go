/*******************************************************************************
实现：日志输出到终端和文件
作者：Lemine
时间：2020/09/09
*******************************************************************************/
package loggo

import (
	"./utils"
)

type ConsoleFileLog struct {
	console *ConsoleLog
	file    *FileLog
}

func NewConsoleFileLog(path, name, runMode string) *ConsoleFileLog {
	if err := checkRunMode(runMode); err != nil {
		panic(err)
	}

	log := new(ConsoleFileLog)

	log.file = NewFileLog(path, name, runMode)
	log.console = NewConsoleLog(runMode)

	return log
}

//------------------------------------------------------------------------------
//Debug等级
func (self *ConsoleFileLog) Debug(message interface{}) {
	if self.file.runMode == RunModeDevelopment && self.file.closeAllChans == false {
		self.console.output.WriteString(utils.StandardMsgWithLineInfo(3, DebugLevel, message))
		self.file.debugChan <- utils.StandardMsgWithLineInfo(3, DebugLevel, message)
	}
}

func (self *ConsoleFileLog) FormatDebug(format string, args ...interface{}) {
	if self.file.runMode == RunModeDevelopment && self.file.closeAllChans == false {
		self.console.output.WriteString(utils.FormatMsgWithLineInfo(4, DebugLevel, format, args...))
		self.file.debugChan <- utils.FormatMsgWithLineInfo(4, DebugLevel, format, args...)
	}
}

//------------------------------------------------------------------------------
//Trace等级
func (self *ConsoleFileLog) Trace(message interface{}) {
	if self.file.runMode == RunModeDevelopment && self.file.closeAllChans == false {
		self.console.output.WriteString(utils.StandardMsgWithLineInfo(3, TraceLevel, message))
		self.file.debugChan <- utils.StandardMsgWithLineInfo(3, TraceLevel, message)
	}
}

func (self *ConsoleFileLog) FormatTrace(format string, args ...interface{}) {
	if self.file.runMode == RunModeDevelopment && self.file.closeAllChans == false {
		self.console.output.WriteString(utils.FormatMsgWithLineInfo(4, TraceLevel, format, args...))
		self.file.debugChan <- utils.FormatMsgWithLineInfo(4, TraceLevel, format, args...)
	}
}

//------------------------------------------------------------------------------
//Info等级
func (self *ConsoleFileLog) Info(message interface{}) {
	if self.file.closeAllChans == false {
		self.console.Info(message)
		self.file.Info(message)
	}
}

func (self *ConsoleFileLog) FormatInfo(format string, args ...interface{}) {
	if self.file.closeAllChans == false {
		self.console.FormatInfo(format, args...)
		self.file.FormatInfo(format, args...)
	}
}

//------------------------------------------------------------------------------
//Warn等级
func (self *ConsoleFileLog) Warn(message interface{}) {
	if self.file.runMode == RunModeProduction && self.file.closeAllChans == false {
		self.console.output.WriteString(utils.StandardMsgWithLineInfo(3, WarnLevel, message))
		self.file.warnChan <- utils.StandardMsgWithLineInfo(3, WarnLevel, message)
	}
}

func (self *ConsoleFileLog) FormatWarn(format string, args ...interface{}) {
	if self.file.runMode == RunModeProduction && self.file.closeAllChans == false {
		self.console.output.WriteString(utils.FormatMsgWithLineInfo(4, WarnLevel, format, args...))
		self.file.warnChan <- utils.FormatMsgWithLineInfo(4, WarnLevel, format, args...)
	}
}

//------------------------------------------------------------------------------
//Error等级
func (self *ConsoleFileLog) Error(message interface{}) {
	if self.file.runMode == RunModeProduction && self.file.closeAllChans == false {
		self.console.output.WriteString(utils.StandardMsgWithLineInfo(3, ErrorLevel, message))
		self.file.warnChan <- utils.StandardMsgWithLineInfo(3, ErrorLevel, message)
	}
}

func (self *ConsoleFileLog) FormatError(format string, args ...interface{}) {
	if self.file.runMode == RunModeProduction && self.file.closeAllChans == false {
		self.console.output.WriteString(utils.FormatMsgWithLineInfo(4, ErrorLevel, format, args...))
		self.file.warnChan <- utils.FormatMsgWithLineInfo(4, ErrorLevel, format, args...)
	}
}

//------------------------------------------------------------------------------
//Fatal等级
func (self *ConsoleFileLog) Fatal(message interface{}) {
	if self.file.runMode == RunModeProduction && self.file.closeAllChans == false {
		self.console.output.WriteString(utils.StandardMsgWithLineInfo(3, FatalLevel, message))
		self.file.warnChan <- utils.StandardMsgWithLineInfo(3, FatalLevel, message)
	}
}

func (self *ConsoleFileLog) FormatFatal(format string, args ...interface{}) {
	if self.file.runMode == RunModeProduction && self.file.closeAllChans == false {
		self.console.output.WriteString(utils.FormatMsgWithLineInfo(4, FatalLevel, format, args...))
		self.file.warnChan <- utils.FormatMsgWithLineInfo(4, FatalLevel, format, args...)
	}
}

func (self *ConsoleFileLog) Close() {
	self.file.Close()
	self.console.output.Close()
}
