package main

import (
	"fmt"
	"math"
	"strings"
)

func basename(str string) string {
	var base []string
	for _, i := range str {
		s := string(i)
		if s == `/` {
			base = []string{}
		} else {
			base = append(base, s)
		}
	}

	dotIndex := len(base)
	for index, d := range base {
		if d == `.` {
			dotIndex = index
		}
	}

	return strings.Join(base[:dotIndex], "")
}

func main() {
	fmt.Println(basename("a/b/c.go")) // "c"
	fmt.Println(basename("c.d.go"))   // "c.d"
	fmt.Println(basename("abc"))      // "abc"
	fmt.Println(math.Pow(2, 10))
}
