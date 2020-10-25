/*******************************************************************************
实现：全局loggo日志
作者：Lemine
时间：2020/09/09
*******************************************************************************/
package loggo

func NewLog(config Config) Logger {
	var log Logger

	if _, ok := config["logPath"]; !ok {
		panic("missing the logPath")
	} else if _, ok := config["logName"]; !ok {
		panic("missing the logName")
	} else if _, ok := config["runMode"]; !ok {
		panic("missing the runMode")
	} else if _, ok := config["outMode"]; !ok {
		panic("missing the outMode")
	}

	if err := checkRunMode(config["runMode"]); err != nil {
		panic(err)
	}

	switch config["outMode"] {
	case OutputToFile:
		log = NewFileLog(config["logPath"], config["logName"], config["runMode"])
	case OutputToConsole:
		log = NewConsoleLog(config["runMode"])
	case OutputToConsoleAndFile:
		log = NewConsoleFileLog(config["logPath"], config["logName"], config["runMode"])
	default:
		panic("unsupport the outMode: " + config["outMode"])
	}

	return log
}
