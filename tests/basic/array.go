package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

func main() {
	var arr = [...]int{1, 2, 3, 4}
	var arr2 = [...]string{1: "a", 3: "c", 2: "b"}
	var arr3 = [...]int{2: 5}
	var arr4 = [...]int{99: -1}
	var arr5 = [...]int{1, 2, 4, 3}
	var arr6 = []byte{'a', 'b'}
	fmt.Println(strings.Split("ab", ""))
	fmt.Println(arr6)
	fmt.Println(arr == arr5)
	fmt.Println(arr4)
	fmt.Println(arr3)
	fmt.Println(arr2)
	fmt.Println(arr)
	sha256.Sum256([]byte("x"))
}
