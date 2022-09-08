package grpc

import (
	"github.com/cairos9000/go_service/constants"
	grp "github.com/cairos9000/go_service/fibo"
	"github.com/cairos9000/go_service/fibonacci"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

type MessageServer struct{}

func ServerGRPC(addr string) error {
	log.Println("gRPC server starts")
	listener, err := net.Listen(constants.Tcp, addr)
	if err != nil {
		log.Println("Failed to start gRPC server")
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
	defer func() {
		if recoveryMessage := recover(); recoveryMessage != nil {
			log.Println(recoveryMessage)
		}
	}()

	resString := ""
	x, y, parseError := fibonacci.ParseArgs(r.Message, constants.Grpc)
	if parseError != nil {
		return nil, parseError
	}

	res, err := fibonacci.Fibo(x, y)
	if err != nil {
		return nil, err
	}

	for _, i := range res {
		resString += strconv.Itoa(i) + " "
	}

	response := &grp.Response{
		Message: resString,
	}

	return response, nil
}
