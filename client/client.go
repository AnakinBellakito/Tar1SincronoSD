package main

import (
	"client/proto"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}
	serviceClient := proto.NewGreeterClient(conn)

	res, err := serviceClient.Notify(context.Background(), &proto.EmptyRequest{})

	if err != nil {
		panic(res)
	}

	fmt.Println(res.GetResponse())
}
