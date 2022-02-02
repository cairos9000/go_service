package main

import (
	"fmt"
	"os"
	"sync"
	"test/fibonacci"
	"test/grpc"
	"test/http"
)

func main()  {
	var err error

	if len(os.Args) != 4{
		fmt.Println("Must be 3 argument - HTTP-server, gRPC-server, Redis")
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)
	httpAddress := os.Args[1]
	gRPCAddress := os.Args[2]
	fmt.Println("HTTP-server address is", httpAddress)
	fmt.Println("gRPC-server address is", gRPCAddress)

	fibonacci.RedisConnection, err = fibonacci.ConnectToRedis(os.Args[3])
	if err != nil{
		fmt.Println(err)
	}

	go func() {
		defer wg.Done()
		err := http.ServerHTTP(httpAddress)
		if err != nil{
			fmt.Println(err)
		}
	}()

	go func() {
		defer wg.Done()
		err := grpc.ServerGRPC(gRPCAddress)
		if err != nil{
			fmt.Println(err)
		}
	}()

	wg.Wait()
	if fibonacci.RedisConnection != nil{
		err = fibonacci.DisconnectRedis(fibonacci.RedisConnection)
		if err != nil{
			fmt.Println(err)
		}
	}
}

