package main

import "fmt"

func main() {
    fmt.Println("main")
    fmt.Println("master chang main function")
}

func foo() {
    fmt.Println("foo")
}

func bar() {
    fmt.Println("bar")
    fmt.Println("dev chang bar function")
}