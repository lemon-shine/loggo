/*******************************************************************************
实现：loggo的基本常量
作者：Lemine
时间：2020/09/09
*******************************************************************************/
package loggo

import (
	"errors"
)

type Config map[string]string

//日志等级
const (
	DebugLevel = "DEBUG" //Debug等级
	TraceLevel = "TRACE" //Trace等级
	InfoLevel  = "INFO"  //Info等级
	WarnLevel  = "WARN"  //Warn等级
	ErrorLevel = "ERROR" //Error等级
	FatalLevel = "FATAL" //Fatal等级
)

func checkLogLevel(level string) error {
	if level == DebugLevel {
		return nil
	} else if level == TraceLevel {
		return nil
	} else if level == InfoLevel {
		return nil
	} else if level == WarnLevel {
		return nil
	} else if level == ErrorLevel {
		return nil
	} else if level == FatalLevel {
		return nil
	}

	return errors.New("unsupport the log level: " + level)
}

//------------------------------------------------------------------------------
//日志调用模式
const (
	RunModeDevelopment = "DEVElOPMENT" //开发模式
	RunModeProduction  = "PRODUCTION"  //生产模式
)

func checkRunMode(mode string) error {
	if mode == RunModeDevelopment {
		return nil
	} else if mode == RunModeProduction {
		return nil
	}

	return errors.New("unsupport the log running mode: " + mode)
}

//------------------------------------------------------------------------------
//日志输出模式
const (
	OutputToFile           = "ONLY_FILE"    //只文件输出
	OutputToConsole        = "ONLY_CONSOLE" //只终端输出
	OutputToConsoleAndFile = "FILE_CONSOLE" //文件和终端同时输出
)

func checkOutputMode(mode string) error {
	if mode == OutputToFile {
		return nil
	} else if mode == OutputToConsole {
		return nil
	} else if mode == OutputToConsoleAndFile {
		return nil
	}

	return errors.New("unsupport the log output mode: " + mode)
}
