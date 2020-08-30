package main

import "fmt"

func main() {
    fmt.Println("main dev change main again")
}

func foo() {
    fmt.Println("foo")
}

func bar() {
    fmt.Println("bar")
    fmt.Println("dev chang bar function")
}