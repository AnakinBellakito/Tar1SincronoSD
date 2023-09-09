package main

import (
	"context"
	"fmt"

	"client4/proto"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}
	serviceClient := proto.NewGreeterClient(conn)

	res, err := serviceClient.QuantKeys(context.Background(), &proto.EmptyRequest{})

	if err != nil {
		fmt.Println(err.Error)
		panic(res)
	}

	fmt.Println(res.GetResponse())
}
