package common

import (
	"github.com/henosteven/heigo/config"
	"log"
	"os"
	"fmt"
)

const (
	PREFIX_INFO = "[INFO]"
	PREFIX_WARNING = "[WARNING]"
	PREFIX_ERROR = "[ERROR]"
)

var OutputMap map[string]*os.File

func InitLog(logpath config.LogPath) {
	log.SetFlags(log.LstdFlags|log.Lshortfile)

	ftrace, err := os.OpenFile(logpath.TracePath, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	OutputMap = make(map[string]*os.File)
	OutputMap[PREFIX_INFO] = ftrace

	ferror, err := os.OpenFile(logpath.ErrorPath, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	OutputMap[PREFIX_WARNING] = ferror
	OutputMap[PREFIX_ERROR] = ferror
}

func LogTrace(trace HeiTrace, msg string) {
	log.SetOutput(OutputMap[PREFIX_INFO])
	logData(trace, PREFIX_INFO, msg)
}

func LogFatal(trace HeiTrace, msg string) {
	log.SetOutput(OutputMap[PREFIX_ERROR])
	logData(trace, PREFIX_ERROR, msg)
}

func LogWarning(trace HeiTrace, msg string) {
	log.SetOutput(OutputMap[PREFIX_WARNING])
	logData(trace, PREFIX_WARNING, msg)
}

func logData(trace HeiTrace, prefix string, msg string) {
	log.SetPrefix(prefix)
	logStr := fmt.Sprintf("%s|msg:%s", trace.GetTraceString(), msg)
	log.Println(logStr)
}