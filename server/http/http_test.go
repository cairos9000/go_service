package http

import (
	"testing"
)

func TestServerHTTP(t *testing.T) {
	e := ServerHTTP("1111")
	if e == nil{
		t.Error("Expected error")
	}

}

