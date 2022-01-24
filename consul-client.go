package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc_consul/pb"
	"strconv"
)

func main() {
	consulConfig := api.DefaultConfig()
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		fmt.Println("Client, api.NewClient err",err)
		return
	}
	services, _, err := consulClient.Health().Service("wbw001_grpc_consul","wbw1",true,nil)
	if err != nil {
		return
	}
	addr := services[0].Service.Address+":"+strconv.Itoa(services[0].Service.Port)
	//grpcConn,err := grpc.Dial("127.0.0.1:8800",grpc.WithInsecure())
	grpcConn,err := grpc.Dial(addr,grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		fmt.Println("dial err:",err)
	}
	defer grpcConn.Close()

	grpcClient := pb.NewHelloClient(grpcConn)

	var person pb.Person
	person.Name = "Lixiaoming"
	person.Age = 18
	per, err := grpcClient.SayHello(context.TODO(), &person)
	if err != nil {
		return
	}
	fmt.Println(per,err)
}