package main

import "fmt"

type human interface {
    sayHi() string
}

type man struct {
    greeting string
}

type woman struct {
    greeting string
}

func (m man) sayHi() string {
    return m.greeting
}

func (w woman) sayHi() string {
    return w.greeting
}

func printGreeting(h human) string {
    return h.sayHi()
}

func main() {
    var m = man{greeting: "I man"}
    var w = woman{greeting: "I woman"}

    fmt.Println(printGreeting(m))
    fmt.Println(printGreeting(w))
}
