package grpc

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"strconv"
	"test/constants"
	grp "test/fibo"
	"test/fibonacci"
)

type MessageServer struct {}


func ServerGRPC(addr string)  error{
	fmt.Println("gRPC server starts")
	listener, err := net.Listen(constants.Tcp, addr)
	if err != nil{
		fmt.Println("Failed to start gRPC server")
		return err
	}

	grpcServer := grpc.NewServer()
	grp.RegisterMessageServiceServer(grpcServer, &MessageServer{})


	if err := grpcServer.Serve(listener); err != nil {
		return err
	}
	return nil
}

func (m *MessageServer) Calc(_ context.Context, r *grp.Request) (*grp.Response, error) {
	defer func(){
		if recoveryMessage:=recover(); recoveryMessage != nil {
			fmt.Println(recoveryMessage)
		}
	}()

	resString := ""
	x, y, parseError := fibonacci.ParseArgs(r.Message, constants.Grpc)
	if parseError != nil{
		return nil, parseError
	}

	res, err := fibonacci.Fibo(x, y)
	if err != nil{
		return nil, err
	}

	for _, i := range res{
		resString += strconv.Itoa(i) + " "
	}

	response := &grp.Response{
		Message: resString,
	}

	return response, nil
}
