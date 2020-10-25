/*******************************************************************************
实现：辅助函数
作者：Lemine
时间：2020/09/09
*******************************************************************************/
package utils

import (
	"runtime"
)

//GetLineInfo：获取DEBUG日志所在的文件、函数名、行号等
func GetLineInfo(skip int) (fileName, funcName string, lineNo int) {
	if pc, file, line, ok := runtime.Caller(skip); ok {
		fileName = file
		funcName = runtime.FuncForPC(pc).Name()
		lineNo = line
	}

	return
}
