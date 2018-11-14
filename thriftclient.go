package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/henosteven/heigo/heiThrift"
	"github.com/henosteven/heigo/config"
	"net"
	"fmt"
	"log"
)

func main()  {
	tSocket, err := thrift.NewTSocket(net.JoinHostPort(config.HOST, config.THRIFT_PORT))
	if err != nil {
		log.Fatalln("tSocket error:", err)
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	transport := transportFactory.GetTransport(tSocket)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	client := heiThrift.NewFormatDataClientFactory(transport, protocolFactory)

	if err := transport.Open(); err != nil {
		log.Fatalln("Error opening:", net.JoinHostPort(config.HOST, config.THRIFT_PORT))
	}
	defer transport.Close()


	data := heiThrift.Data{Text:"hello,world!"}
	d, err := client.DoFormat(&data)
	fmt.Println(d.Text)
}