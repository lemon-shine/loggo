package loggo

import (
	"testing"
	"time"
)

// func TestDevelopmentConsoleLog(t *testing.T) {
// 	log := NewConsoleFileLog("./", "test", RunModeDevelopment)

// 	go func() {
// 		for i := 0; i < 100; i++ {
// 			log.Debug("this is a debug log")
// 			log.FormatDebug("this is a %s log %d", "debug", i)
// 			log.Trace("this is a trace log")
// 			log.FormatTrace("this is a %s log %d", "trace", i)
// 		}
// 	}()

// 	time.Sleep(10 * time.Second)
// 	t.Log("test finished")
// 	log.Close()
// }

func TestProductionConsoleLog(t *testing.T) {
	log := NewConsoleFileLog("./", "test", RunModeProduction)

	go func() {
		for i := 0; i < 100; i++ {
			log.Info("this is a info log")
			log.FormatInfo("this is a %s log %d", "info", i)
			log.Warn("this is a warn log")
			log.FormatWarn("this is a %s log %d", "warn", i)
			log.Error("this is a error log")
			log.FormatError("this is a %s log %d", "error", i)
			log.Fatal("this is a fatal log")
			log.FormatFatal("this is a %s log %d", "fatal", i)
		}
	}()

	time.Sleep(10 * time.Second)
	t.Log("test finished")
	log.Close()
}
