package main

import (
	"context"
	"flag"
	"fmt"
	"strconv"
	"time"

	"google.golang.org/grpc"

	pb "github.com/panyingyun/grpcglb/cmd/helloworld"
	grpclb "github.com/panyingyun/grpcglb/etcdv3"
)

var (
	serv = flag.String("service", "grpcglb", "service name")
	reg  = flag.String("reg", "http://123.206.185.178:2379", "register etcd address")
)

func main() {
	flag.Parse()
	r := grpclb.NewResolver(*serv)

	b := grpc.RoundRobin(r)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	conn, err := grpc.DialContext(ctx, *reg, grpc.WithInsecure(), grpc.WithBalancer(b), grpc.WithBlock())
	cancel()
	if err != nil {
		panic(err)
	}

	ticker := time.NewTicker(time.Second)
	for t := range ticker.C {
		client := pb.NewGreeterClient(conn)
		resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "world " + strconv.Itoa(t.Second())})
		if err == nil {
			fmt.Printf("%v: Reply is %s\n", t.Local(), resp.Message)
		}
	}
}
