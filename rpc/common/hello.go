package common

import "net/rpc"

/*
服务端、客户端两端的接口定义
*/

const HelloServiceName = "awesome-go/rpc.HelloService"

type IHelloService = interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(impl IHelloService) error {
	return rpc.RegisterName(HelloServiceName, impl)
}

type HelloServiceClient struct {
	*rpc.Client
}

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}
