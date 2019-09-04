package main

import (
	r "awesome-go/rpc"
	"awesome-go/rpc/common"
	"io"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//
//func main() {
//    if err := rpc.RegisterName("HelloService", new(r.HelloService)); err != nil {
//        log.Fatal("Regist service error:", err)
//    }
//
//    listener, err := net.Listen("tcp", ":7000")
//    if err != nil {
//        log.Fatal("ListenTCP error:", err)
//    }
//
//    conn, err := listener.Accept()
//    if err != nil {
//        log.Fatal("Accept error:", err)
//    }
//
//    rpc.ServeConn(conn)
//}

// 封装后的服务注册
//func main() {
//    err := common.RegisterHelloService(new(r.HelloService))
//    if err != nil {
//        log.Fatal("register helloService error: ", err)
//    }
//
//    listener, err := net.Listen("tcp", ":7000")
//    if err != nil {
//       log.Fatal("ListenTCP error:", err)
//    }
//    for {
//        conn, err := listener.Accept()
//        if err != nil {
//            log.Fatal("accept error:", err)
//        }
//
//        go rpc.ServeConn(conn)
//    }
//}

// 跨语言rpc，json
//func main() {
//    err := common.RegisterHelloService(new(r.HelloService))
//    if err != nil {
//        log.Fatal("register helloService error: ", err)
//    }
//
//    listener, err := net.Listen("tcp", ":7000")
//    if err != nil {
//        log.Fatal("ListenTCP error:", err)
//    }
//    for {
//        conn, err := listener.Accept()
//        if err != nil {
//            log.Fatal("accept error:", err)
//        }
//
//        go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
//    }
//}

// 基于http的json rpc
func main() {
	err := common.RegisterHelloService(new(r.HelloService))
	if err != nil {
		log.Fatal("register helloService error: ", err)
	}

	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}

		err := rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
		if err != nil {
			log.Println("server error:", err)
		}
	})

	err = http.ListenAndServe(":7000", nil)
	if err != nil {
		log.Fatal("listen error:", err)
	}
}
