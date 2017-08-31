package main

import (
	"fmt"
	"github.com/timeloveboy/taskmaster/pb"
	"google.golang.org/grpc"

	"golang.org/x/net/context"
)

func main() {
	conn, err := grpc.Dial("localhost:10000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)
	result, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "李鹏"})
	fmt.Println(result.Message, err)
}
