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
	client := pb.NewTaskMasterClient(conn)
	connectcli, err := client.Connect(context.Background())
	for {
		connectcli.Send(&pb.OwnerState{
			Owner:  "me",
			TaskId: 1024,
			Group:  "1024",
		})
		tr, err := connectcli.Recv()
		fmt.Println(tr)

	}

}
