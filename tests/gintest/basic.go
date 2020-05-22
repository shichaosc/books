package main

import (
	"fmt"
	"runtime"
	"strings"
)

func main() {
	versionTag := runtime.Version()
	fmt.Println(versionTag)
	first := strings.IndexByte(runtime.Version(), '.')
	fmt.Println(first)
	last := strings.LastIndexByte(runtime.Version(), '.')
	fmt.Println(last)
	fmt.Println(versionTag[first+1 : last])
}
