package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/codegangsta/martini"
	"github.com/henosteven/heigo/httpservice"
	"github.com/henosteven/heigo/heiThrift"
	"fmt"
	"runtime"
	"os"
	"os/signal"
	"syscall"
	"time"
	"log"
)

var quit = make(chan int)

type FormatDataImpl struct{}

func (fdi FormatDataImpl) DoFormat (data *heiThrift.Data) (r *heiThrift.Data, err error) {
	return data, nil
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	go signalProcess()
	go initMartini()
	go initThriftServe()
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

func initThriftServe() {
	handler := &FormatDataImpl{}
	processor := heiThrift.NewFormatDataProcessor(handler)
	serverTransport, err := thrift.NewTServerSocket("127.0.0.1:3001")
	if err != nil {
		log.Fatalln("Error:", err)
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("Running at:", "127.0.0.1:3001")
	server.Serve()
}
