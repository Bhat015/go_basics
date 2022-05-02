package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")

	// response body
	// reader and closer as interface
	// anything which has a Read() and Close() method can be response

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))

	// copy the response body
	io.Copy(os.Stdout, resp.Body)
}
