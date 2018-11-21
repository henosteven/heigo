# heigo
测试使用功能的小项目

## thrift文件生成
thrift --out ../  --gen go hei.thrift


```
package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/henosteven/heigo/heiThrift"
	"net"
	"fmt"
	"log"
	"context"
)

const (
	HOST = "localhost"
	PORT = "3000"
)

func main()  {
	tSocket, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
	if err != nil {
		log.Fatalln("tSocket error:", err)
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	transport, _:= transportFactory.GetTransport(tSocket)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	client := heiThrift.NewUserHandlerClientFactory(transport, protocolFactory)

	if err := transport.Open(); err != nil {
		log.Fatalln("Error opening:", HOST + ":" + PORT)
	}
	defer transport.Close()

	ctx := context.Background()
	user, err := client.GetUser(ctx, 1)
	fmt.Println(user, err)

	userID, err := client.AddUser(ctx, "haoran")
	fmt.Println(userID, err)
}
```