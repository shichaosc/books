package main

import "fmt"

type greeting func(s string) string

func (g greeting) say(n string) {
	g(n)
}

func english(s string) string {
	fmt.Println(s)
	return s
}

func chinese(s string) string {
	fmt.Printf("%s chinese", s)
	return s
}

func main() {
	g := greeting(chinese)
	g.say("s")
}
