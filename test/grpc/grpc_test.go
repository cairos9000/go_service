package grpc

import "testing"

func TestServerGRPC(t *testing.T) {
	err := ServerGRPC("1111")
	if err == nil{
		t.Error("Expected error")
	}
	err = ServerGRPC("1.1.1.1")
	if err == nil{
		t.Error("Expected error")
	}

}
