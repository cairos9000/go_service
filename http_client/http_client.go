package main

import (
	"fmt"
	"io"
	"net/http"
)

func httpClient(){
	resp, err := http.Get("http://127.0.0.1:9000/[30;40]")
	if err != nil{
		fmt.Println(err.Error())
	} else {
		defer resp.Body.Close()
		if resp.StatusCode == 444{
			fmt.Println("Request failed with error. Code status is 444")
			body, _ := io.ReadAll(resp.Body)
			fmt.Println(string(body))

		} else{
			body, _ := io.ReadAll(resp.Body)
			fmt.Println(string(body))
		}

	}
}

func main()  {
	httpClient()


}
