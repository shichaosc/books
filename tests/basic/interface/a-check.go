package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	var a io.Writer
	i, ok := a.(io.Writer)
	fmt.Println(ok)
	fmt.Println(i)
	strings.NewReader()
}
