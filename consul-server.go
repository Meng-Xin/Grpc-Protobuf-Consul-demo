package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"grpc_consul/pb"
	"net"
)

type PersonServer struct {
	pb.UnimplementedHelloServer
}

func (ps *PersonServer)SayHello(ctx context.Context,person *pb.Person)(*pb.Person,error){
	person.Name = "hello-" + person.Name
	return person,nil
}

func main()  {
	//把grpc服务注册到consul上
	//初始化consul配置
	consulConfig := api.DefaultConfig()
	//创建consul对象
	consulClient,err := api.NewClient(consulConfig)
	if err!=nil{
		fmt.Println("server, api.newclient err:",err)
		return
	}
	//告诉consul，即将注册的服务的配置信息
	reg := api.AgentServiceRegistration{
		Kind:              "",
		ID:                "wbw1_id",
		Name:              "wbw001_grpc_consul",
		Tags:              []string{"wbw1","asa1"},
		Port:              8800,
		Address:           "127.0.0.1",
		TaggedAddresses:   nil,
		EnableTagOverride: false,
		Meta:              nil,
		Weights:           nil,
		Check:             &api.AgentServiceCheck{
			CheckID:                        "wbw1_id_check",
			Name:                           "",
			Args:                           nil,
			DockerContainerID:              "",
			Shell:                          "",
			Interval:                       "5s",
			Timeout:                        "1s",
			TTL:                            "",
			HTTP:                           "",
			Header:                         nil,
			Method:                         "",
			Body:                           "",
			TCP:                            "127.0.0.1:8800",
			Status:                         "",
			Notes:                          "",
			TLSServerName:                  "",
			TLSSkipVerify:                  false,
			GRPC:                           "",
			GRPCUseTLS:                     false,
			AliasNode:                      "",
			AliasService:                   "",
			SuccessBeforePassing:           0,
			FailuresBeforeCritical:         0,
			DeregisterCriticalServiceAfter: "",
		},
		Checks:            nil,
		Proxy:             nil,
		Connect:           nil,
		Namespace:         "",
	}
	//注册gprc服务到consul上
	consulClient.Agent().ServiceRegister(&reg)
	grpcServer := grpc.NewServer()
	pb.RegisterHelloServer(grpcServer,&PersonServer{})
	listener,err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil{
		fmt.Println("listen err:",err)
	}
	defer listener.Close()
	grpcServer.Serve(listener)

}