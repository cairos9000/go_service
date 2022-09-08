package main

import (
	"fmt"
	"github.com/cairos9000/go_service/grpc"
	"github.com/cairos9000/go_service/http"
	"github.com/cairos9000/go_service/redis"
	"log"
	"os"
	"sync"
)

func StartGRPC(address string) {
	err := grpc.ServerGRPC(address)
	if err != nil {
		fmt.Println(err)
	}
}

func StartHTPP(address string) {
	err := http.ServerHTTP(address)
	if err != nil {
		fmt.Println(err)
	}
}

func StartServer(serverType string, address string) {
	switch serverType {
	case "HTTP":
		StartHTPP(address)
	case "gRPC":
		StartGRPC(address)
	}
}

func main() {
	var err error
	var wg sync.WaitGroup

	if len(os.Args) != 4 {
		log.Println("Must be 3 argument - HTTP-server, gRPC-server, Redis")
		return
	}

	httpAddress := os.Args[1]
	gRPCAddress := os.Args[2]
	redisAddress := os.Args[3]
	servers := []string{"HTTP", "gRPC"}
	addresses := []string{httpAddress, gRPCAddress}

	log.Println("HTTP-server address is", httpAddress)
	log.Println("gRPC-server address is", gRPCAddress)

	redis.Connection, err = redis.ConnectToRedis(redisAddress)
	if err != nil {
		log.Println(err)
	}

	wg.Add(len(servers))
	for ind, s := range servers {
		go func(serverType string, address string) {
			defer wg.Done()
			StartServer(serverType, address)
		}(s, addresses[ind])
	}

	wg.Wait()
	if redis.Connection != nil {
		err = redis.DisconnectRedis(redis.Connection)
		if err != nil {
			log.Println(err)
		}
	}
}
