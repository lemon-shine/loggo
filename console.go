/*******************************************************************************
实现：日志输出到终端
作者：Lemine
时间：2020/09/09
*******************************************************************************/
package loggo

import (
	"os"

	"./utils"
)

//定义终端打印日志
type ConsoleLog struct {
	output  *os.File
	runMode string
}

func NewConsoleLog(runMode string) *ConsoleLog {
	if err := checkRunMode(runMode); err != nil {
		panic(err)
	}

	return &ConsoleLog{
		output:  os.Stdout,
		runMode: runMode,
	}
}

//------------------------------------------------------------------------------
//Debug等级
func (self *ConsoleLog) Debug(message interface{}) {
	if self.runMode == RunModeDevelopment {
		self.output.WriteString(utils.StandardMsgWithLineInfo(3, DebugLevel, message))
	}
}

func (self *ConsoleLog) FormatDebug(format string, args ...interface{}) {
	if self.runMode == RunModeDevelopment {
		self.output.WriteString(utils.FormatMsgWithLineInfo(4, DebugLevel, format, args...))
	}
}

//------------------------------------------------------------------------------
//Trace等级
func (self *ConsoleLog) Trace(message interface{}) {
	if self.runMode == RunModeDevelopment {
		self.output.WriteString(utils.StandardMsgWithLineInfo(3, TraceLevel, message))
	}
}

func (self *ConsoleLog) FormatTrace(format string, args ...interface{}) {
	if self.runMode == RunModeDevelopment {
		self.output.WriteString(utils.FormatMsgWithLineInfo(4, TraceLevel, format, args...))
	}
}

//------------------------------------------------------------------------------
//Info等级
func (self *ConsoleLog) Info(message interface{}) {
	if self.runMode == RunModeProduction {
		self.output.WriteString(utils.StandardMsgWithoutLineInfo(InfoLevel, message))
	}
}

func (self *ConsoleLog) FormatInfo(format string, args ...interface{}) {
	if self.runMode == RunModeProduction {
		self.output.WriteString(utils.FormatMsgWithoutLineInfo(InfoLevel, format, args...))
	}
}

//------------------------------------------------------------------------------
//Warn等级
func (self *ConsoleLog) Warn(message interface{}) {
	if self.runMode == RunModeProduction {
		self.output.WriteString(utils.StandardMsgWithLineInfo(3, WarnLevel, message))
	}
}

func (self *ConsoleLog) FormatWarn(format string, args ...interface{}) {
	if self.runMode == RunModeProduction {
		self.output.WriteString(utils.FormatMsgWithLineInfo(4, WarnLevel, format, args...))
	}
}

//------------------------------------------------------------------------------
//Error等级
func (self *ConsoleLog) Error(message interface{}) {
	if self.runMode == RunModeProduction {
		self.output.WriteString(utils.StandardMsgWithLineInfo(3, ErrorLevel, message))
	}
}

func (self *ConsoleLog) FormatError(format string, args ...interface{}) {
	if self.runMode == RunModeProduction {
		self.output.WriteString(utils.FormatMsgWithLineInfo(4, ErrorLevel, format, args...))
	}
}

//------------------------------------------------------------------------------
//Fatal等级
func (self *ConsoleLog) Fatal(message interface{}) {
	if self.runMode == RunModeProduction {
		self.output.WriteString(utils.StandardMsgWithLineInfo(3, ErrorLevel, message))
	}
}

func (self *ConsoleLog) FormatFatal(format string, args ...interface{}) {
	if self.runMode == RunModeProduction {
		self.output.WriteString(utils.FormatMsgWithLineInfo(4, ErrorLevel, format, args...))
	}
}

func (self *ConsoleLog) Close() {
	self.output.Close()
}
