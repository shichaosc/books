package main

import (
	"fmt"
	"net/http"
)

func main() {
	serve := http.ListenAndServe("127.0.0.1:8888", nil)
	fmt.Println(serve)
}
