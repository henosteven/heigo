package main

import (
	"github.com/codegangsta/martini"
	"github.com/henosteven/heigo/httpservice"
	"fmt"
	"runtime"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var quit = make(chan int)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	go signalProcess()
	go initMartini()
	<- quit
	fmt.Println("ctrl -c ~ bye~bye~")
	time.Sleep(time.Second * 2)
}

func signalProcess() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	signal.Notify(c, syscall.SIGINT)
	<-c
	quit <- 1
}

func initMartini() {
	m := martini.Classic()
	m.Get("/", httpservice.Hello)
	m.Get("/get", httpservice.Get)
	m.Get("/set", httpservice.Set)
	m.Run()
}
