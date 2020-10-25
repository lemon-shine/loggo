/*******************************************************************************
实现：日志输出到文件
作者：Lemine
时间：2020/09/09
*******************************************************************************/
package loggo

import (
	"fmt"
	"os"

	"./utils"
)

const (
	defaultLen = 20000
)

//定义文件日志结构
type FileLog struct {
	logPath       string //日志文件路径
	logName       string //日志文件名
	chanEnd       string //管道结束标识符
	runMode       string //日志调用模式
	closeAllChans bool   //是否所有管道已关闭

	debugFile *os.File //调试文件日志
	infoFile  *os.File //普通文件日志
	warnFile  *os.File //异常文件日志

	closeChan chan bool   //用于关闭日志管道
	debugChan chan string //用于异步写日志信息到文件
	infoChan  chan string
	warnChan  chan string
}

func NewFileLog(path, name, runMode string) *FileLog {
	if err := checkRunMode(runMode); err != nil {
		panic(err)
	}

	log := &FileLog{
		logPath:   path,
		logName:   name,
		runMode:   runMode,
		chanEnd:   "\r\t\r\t",
		closeChan: make(chan bool),
	}

	if runMode == RunModeDevelopment {
		log.debugChan = make(chan string, defaultLen)
	}

	if runMode == RunModeProduction {
		log.infoChan = make(chan string, defaultLen)
		log.warnChan = make(chan string, defaultLen)
	}

	log.init()
	go log.writeLogToFile()

	return log
}

func (self *FileLog) init() {
	if self.runMode == RunModeDevelopment {
		//创建调试日志文件
		debugFilename := fmt.Sprintf("%s%s.debug.log", self.logPath, self.logName)
		debugFile, err := os.OpenFile(debugFilename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
		if err != nil {
			panic(err)
		}
		self.debugFile = debugFile

	} else if self.runMode == RunModeProduction {
		//创建普通日志文件
		infoFilename := fmt.Sprintf("%s%s.info.log", self.logPath, self.logName)
		infoFile, err := os.OpenFile(infoFilename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
		if err != nil {
			panic(err)
		}
		self.infoFile = infoFile

		//创建异常日志文件
		warnFileName := fmt.Sprintf("%s%s.warn.log", self.logPath, self.logName)
		warnFile, err := os.OpenFile(warnFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
		if err != nil {
			panic(err)
		}
		self.warnFile = warnFile
	}
}

func (self *FileLog) writeLogToFile() {
	var allChansClosed int //如果为2，表示所有文件和管道都已关闭
	for {
		if self.runMode == RunModeDevelopment {
			select {

			case debugData := <-self.debugChan:
				//若接收到chanEnd，表示debugChan已关闭，则立即关闭文件，并结束协程
				if debugData == self.chanEnd {
					self.debugFile.WriteString("\nThe Debug log is closed normally\n")
					self.debugFile.Close()
					return
				}
				self.debugFile.WriteString(debugData)

			case isClose := <-self.closeChan:
				//关闭debugChan
				if isClose == true {
					self.debugChan <- self.chanEnd //先写入结束标志到debugChan
					close(self.debugChan)          //再关闭debugChan
				}
			}
		}

		if self.runMode == RunModeProduction {
			//若所有的infoFile与infoChan、warnFile与warnChan都已关闭，则立即结束协程
			if allChansClosed == 2 {
				self.closeAllChans = true
				return
			}
			select {

			case infoData := <-self.infoChan:
				//若接收到chanEnd，表示infoChan已关闭，则立即关闭infoFile文件
				if infoData != self.chanEnd {
					self.infoFile.WriteString(infoData)
				} else {
					self.infoFile.WriteString("\nThe Info log is closed normally\n")
					self.infoFile.Close()
					allChansClosed++
				}

			case warnData := <-self.warnChan:
				//若接收到chanEnd，表示warnChan已关闭，则立即关闭warnFile文件
				if warnData != self.chanEnd {
					self.warnFile.WriteString(warnData)
				} else {
					self.warnFile.WriteString("\nThe Warn log is closed normally\n")
					self.warnFile.Close()
					allChansClosed++
				}

			case isClose := <-self.closeChan:
				//关闭infoChan和warnChan
				if isClose == true {
					//先写入结束标志到infoChan和warnChan，再关闭这两个管道
					self.infoChan <- self.chanEnd
					close(self.infoChan)
					self.warnChan <- self.chanEnd
					close(self.warnChan)
				}
			}
		}
	}
}

//------------------------------------------------------------------------------
//Debug等级
func (self *FileLog) Debug(message interface{}) {
	if self.runMode == RunModeDevelopment && self.closeAllChans == false {
		self.debugChan <- utils.StandardMsgWithLineInfo(3, DebugLevel, message)
	}
}

func (self *FileLog) FormatDebug(format string, args ...interface{}) {
	if self.runMode == RunModeDevelopment && self.closeAllChans == false {
		self.debugChan <- utils.FormatMsgWithLineInfo(4, DebugLevel, format, args...)
	}
}

//------------------------------------------------------------------------------
//Trace等级
func (self *FileLog) Trace(message interface{}) {
	if self.runMode == RunModeDevelopment && self.closeAllChans == false {
		self.debugChan <- utils.StandardMsgWithLineInfo(3, TraceLevel, message)
	}
}

func (self *FileLog) FormatTrace(format string, args ...interface{}) {
	if self.runMode == RunModeDevelopment && self.closeAllChans == false {
		self.debugChan <- utils.FormatMsgWithLineInfo(4, TraceLevel, format, args...)
	}
}

//------------------------------------------------------------------------------
//Info等级
func (self *FileLog) Info(message interface{}) {
	if self.runMode == RunModeProduction && self.closeAllChans == false {
		self.infoChan <- utils.StandardMsgWithoutLineInfo(InfoLevel, message)
	}
}

func (self *FileLog) FormatInfo(format string, args ...interface{}) {
	if self.runMode == RunModeProduction && self.closeAllChans == false {
		self.infoChan <- utils.FormatMsgWithoutLineInfo(InfoLevel, format, args...)
	}
}

//------------------------------------------------------------------------------
//Warn等级
func (self *FileLog) Warn(message interface{}) {
	if self.runMode == RunModeProduction && self.closeAllChans == false {
		self.warnChan <- utils.StandardMsgWithLineInfo(3, WarnLevel, message)
	}
}

func (self *FileLog) FormatWarn(format string, args ...interface{}) {
	if self.runMode == RunModeProduction && self.closeAllChans == false {
		self.warnChan <- utils.FormatMsgWithLineInfo(4, WarnLevel, format, args...)
	}
}

//------------------------------------------------------------------------------
//Error等级
func (self *FileLog) Error(message interface{}) {
	if self.runMode == RunModeProduction && self.closeAllChans == false {
		self.warnChan <- utils.StandardMsgWithLineInfo(3, ErrorLevel, message)
	}
}

func (self *FileLog) FormatError(format string, args ...interface{}) {
	if self.runMode == RunModeProduction && self.closeAllChans == false {
		self.warnChan <- utils.FormatMsgWithLineInfo(4, ErrorLevel, format, args...)
	}
}

//------------------------------------------------------------------------------
//Fatal等级
func (self *FileLog) Fatal(message interface{}) {
	if self.runMode == RunModeProduction && self.closeAllChans == false {
		self.warnChan <- utils.StandardMsgWithLineInfo(3, FatalLevel, message)
	}
}

func (self *FileLog) FormatFatal(format string, args ...interface{}) {
	if self.runMode == RunModeProduction && self.closeAllChans == false {
		self.warnChan <- utils.FormatMsgWithLineInfo(4, FatalLevel, format, args...)
	}
}

func (self *FileLog) Close() {
	self.closeAllChans = true //将该标志位置为true，以阻止日志继续写入
	self.closeChan <- true    //发送关闭信号后，关闭该管道
	close(self.closeChan)
}
