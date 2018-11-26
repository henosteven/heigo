package main

import (
	"context"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/henosteven/heigo/config"
	"github.com/henosteven/heigo/heiThrift"
	"log"
	"net"
)

func main() {
	tSocket, err := thrift.NewTSocket(net.JoinHostPort(config.HOST, config.THRIFT_PORT))
	if err != nil {
		log.Fatalln("tSocket error:", err)
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	transport, _ := transportFactory.GetTransport(tSocket)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	client := heiThrift.NewFormatDataClientFactory(transport, protocolFactory)

	if err := transport.Open(); err != nil {
		log.Fatalln("Error opening:", net.JoinHostPort(config.HOST, config.THRIFT_PORT))
	}
	defer transport.Close()

	ctx := context.Background()
	data := heiThrift.Data{Text: "hello,world!"}
	d, err := client.DoFormat(ctx, &data)
	fmt.Println(d.Text)
}
