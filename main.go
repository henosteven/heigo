package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/codegangsta/martini"
	"github.com/henosteven/heigo/httpservice"
	"github.com/henosteven/heigo/heiThrift"
	"github.com/henosteven/heigo/config"
	"github.com/henosteven/heigo/thriftservice"
	"github.com/henosteven/heigo/model"
	"github.com/henosteven/heigo/common"
	"github.com/henosteven/heigo/lib"
	"fmt"
	"runtime"
	"os"
	"os/signal"
	"syscall"
	"time"
	"log"
	"net"
	"flag"
)

var quit = make(chan int)

const (
	VERSION = "1.0.0"
	DEFAULT_CONFIG_PATH = "./config/conf.toml"
)

var configPath *string

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if handleFlag() {
		os.Exit(0)
	}

	config.InitConfig(*configPath)

	model.InitDb(config.GlobalConfig.MysqlConf)
	lib.InitRedis(config.GlobalConfig.RedisConf)
	common.InitLog(config.GlobalConfig.LogPath)
	common.InitLimitConfig()

	go signalProcess()
	go initMartini()
	go initThriftServe()
	<- quit
	model.TeardownDb()
	fmt.Println("ctrl -c ~ bye~bye~")
	time.Sleep(time.Second * 1)
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
	m.Get("/user", httpservice.SafeHandler(httpservice.GetUser))
	m.Post("/user", httpservice.SafeHandler(httpservice.SetUser))
	os.Setenv("PORT", config.GlobalConfig.WebConf.Port)
	m.Run()
}

func initThriftServe() {
	handler := &thriftservice.UserHandlerImpl{}
	processor := heiThrift.NewUserHandlerProcessor(handler)
	serverTransport, err := thrift.NewTServerSocket(net.JoinHostPort(config.GlobalConfig.Host, config.GlobalConfig.ThriftConf.Port))
	if err != nil {
		log.Fatalln("Error:", err)
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("Running at:", net.JoinHostPort(config.GlobalConfig.Host, config.GlobalConfig.ThriftConf.Port))
	server.Serve()
}

func handleFlag() (needExit bool) {
	needExit = false
	showVersion := flag.Bool("version", false, "version")
	showV := flag.Bool("v", false, "version")

	configPath = flag.String("config", DEFAULT_CONFIG_PATH, "version")

	flag.Parse()
	if *showVersion || *showV {
		fmt.Println("verison:", VERSION)
		needExit = true
	}
	return
}
