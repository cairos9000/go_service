package http

import (
	"fmt"
	"net/http"
	"test/constants"
	"test/fibonacci"
)

func GetHandler(w http.ResponseWriter, req *http.Request) {
	defer func(){
		if recoveryMessage:=recover(); recoveryMessage != nil {
			fmt.Println(recoveryMessage)
		}
	}()

	switch req.Method {
	case constants.HttpGetMethod:
		x, y, parseError := fibonacci.ParseArgs(req.URL.Path, constants.Http)
		if parseError != nil{
			http.Error(w, parseError.Error(), 444)
			break
		}
		res, err := fibonacci.Fibo(x, y)
		if err != nil{
			http.Error(w, err.Error(), 444)
			break
		}
		for _, i := range res{
			_, err = fmt.Fprintf(w, "%d ", i)
			if err != nil{
				http.Error(w, "\nSome error with http server while formatting data", 444)
				break
			}
		}

	default:
		http.Error(w, "This server supports only GET method",444)
	}
}

func ServerHTTP(addr string) error{
	fmt.Println("HTTP server starts")
	http.HandleFunc("/", GetHandler)
	err := http.ListenAndServe(addr, nil)
	if err != nil{
		return err
	}
	return nil
}
