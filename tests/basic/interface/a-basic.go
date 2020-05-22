package main

import "fmt"

type byteCounter int

func (b *byteCounter) Write(p []byte) (int, error) {
	*b += byteCounter(len(p))
	return len(p), nil
}

func main() {
	var b byteCounter
	n, _ := fmt.Fprint(&b, "hello")
	i, _ := fmt.Fprint(&b, "hello1")
	fmt.Println(n)
	fmt.Println(i)
	fmt.Println(b)
}
