package utils

import (
	"fmt"
	"path"
	"time"
)

func StandardMsgWithLineInfo(skip int, level string,
	message interface{}) string {
	timeStr := time.Now().Format("2006/01/02 15:04:05")

	fileName, funcName, lineNo := GetLineInfo(skip)
	fileName = path.Base(fileName)
	funcName = path.Base(funcName)

	return fmt.Sprintf("[%s]\t[%s] %s\n\t\tfile: %s ===> func: %s ===> line: %d\n",
		level, timeStr, message, fileName, funcName, lineNo)
}

func FormatMsgWithLineInfo(skip int, level string,
	format string, args ...interface{}) string {
	msg := fmt.Sprintf(format, args...)
	return StandardMsgWithLineInfo(skip, level, msg)
}

//------------------------------------------------------------------------------

func StandardMsgWithoutLineInfo(level string, message interface{}) string {
	timeStr := time.Now().Format("2006/01/02 15:04:05")

	return fmt.Sprintf("[%s]\t[%s] %s\n", level, timeStr, message)
}

func FormatMsgWithoutLineInfo(level string, format string, args ...interface{}) string {
	msg := fmt.Sprintf(format, args...)
	return StandardMsgWithoutLineInfo(level, msg)
}
