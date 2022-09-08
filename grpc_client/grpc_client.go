package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("127.0.0.1:9001", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	c := NewMessageServiceClient(conn)

	response, err := c.Calc(context.Background(), &Request{Message: "[60;70]"})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.Message)
	}

}
