package common

import (
	"log"
	"os"
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
	log.SetPrefix("INFO")
	log.Println(msg)
}