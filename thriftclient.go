package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/henosteven/heigo/heiThrift"
	"net"
	"fmt"
	"log"
)

const (
	HOST = "localhost"
	PORT = "3001"
)

func main()  {
	tSocket, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
	if err != nil {
		log.Fatalln("tSocket error:", err)
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	transport := transportFactory.GetTransport(tSocket)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	client := heiThrift.NewFormatDataClientFactory(transport, protocolFactory)

	if err := transport.Open(); err != nil {
		log.Fatalln("Error opening:", HOST + ":" + PORT)
	}
	defer transport.Close()


	data := heiThrift.Data{Text:"hello,world!"}
	d, err := client.DoFormat(&data)
	fmt.Println(d.Text)
}