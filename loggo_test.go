package loggo

import (
	"testing"
	"time"
)

func TestLoggo(t *testing.T) {
	log := NewLog(Config{
		"logPath": "./",
		"logName": "hello",
		"runMode": RunModeDevelopment,
		"outMode": OutputToFile,
	})

	t.Log("the development mode is start")
	go func() {
		for i := 0; i < 100; i++ {
			log.Debug("Debug")
			log.FormatDebug("Debug %d", i)
			log.Trace("Trace")
			log.FormatTrace("Trace %d", i)
			log.Info("Info")
			log.FormatInfo("Info %d", i)
			log.Warn("Warn")
			log.FormatWarn("Warn %d", i)
			log.Error("Error")
			log.FormatError("Error %d", i)
			log.Fatal("Fatal")
			log.FormatFatal("Fatal %d", i)
		}

	}()

	time.Sleep(20 * time.Second)
	log.Close()
	t.Log("the development log is start")
}
