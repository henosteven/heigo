package common

import (
	"log"
	"os"
	"fmt"
)

const (
	PREFIX_INFO = "[INFO]"
	PREFIX_WARNING = "[WARNING]"
	PREFIX_ERROR = "[ERROR]"
)

func InitLog(logpath string) {
	log.SetFlags(log.LstdFlags|log.Lshortfile)

	f, err := os.OpenFile(logpath, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(f)
}

func LogTrace(trace HeiTrace, msg string) {
	logData(trace, PREFIX_INFO, msg)
}

func LogFatal(trace HeiTrace, msg string) {
	logData(trace, PREFIX_ERROR, msg)
}

func LogWarning(trace HeiTrace, msg string) {
	logData(trace, PREFIX_WARNING, msg)
}

func logData(trace HeiTrace, prefix string, msg string) {
	log.SetPrefix(prefix)
	logStr := fmt.Sprintf("%s|msg:%s", trace.GetTraceString(), msg)
	log.Println(logStr)
}