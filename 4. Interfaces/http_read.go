package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type customWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
	cw := customWriter{}
	io.Copy(cw, resp.Body)
}

func (cw customWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs), len(bs))
	return len(bs), nil
}
