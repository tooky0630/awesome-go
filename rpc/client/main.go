package main

import (
	"awesome-go/rpc/common"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//
//func main() {
//    client, err := rpc.Dial("tcp", "localhost:7000")
//    if err != nil {
//        log.Fatal("Dial error:", err)
//    }
//
//
//    var reply string
//    // Call()方法 同步调用
//    //err = client.Call("HelloService.Hello", "rpc", &reply)
//    //if err != nil {
//    //    log.Fatal("Call error:", err)
//    //}
//    //
//    //fmt.Println("result: ", reply)
//
//    // Go()，异步调用，从channel获取执行完成状态
//    done := make(chan *rpc.Call, 1)
//    resultFunc := make(map[*rpc.Call]func())
//    resultFunc[client.Go("HelloService.Hello", "rpc", &reply, done)] = func() {
//        fmt.Println("result:", reply)
//    }
//    fmt.Println("now: ", reply)
//    for {
//        select {
//        case call := <- done: // Call from channel == Call returned by Go()
//            resultFunc[call]()
//            return
//        }
//    }
//
//}

// Service封装后的调用
//func main() {
//    helloService, err := common.DialHelloService("tcp", "localhost:7000")
//    if err != nil {
//        log.Fatal("dial HelloService error:", err)
//    }
//
//    var reply string
//    err = helloService.Hello("helloService", &reply)
//    if err != nil {
//        log.Fatal("call servie error:", err)
//    }
//    fmt.Println("result: ", reply)
//}

// 跨语言客户端调用，json
func main() {
	conn, err := net.Dial("tcp", "localhost:7000")
	if err != nil {
		log.Fatal("net.Dial:", err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	err = client.Call(common.HelloServiceName+".Hello", "json", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result: ", reply)
}
