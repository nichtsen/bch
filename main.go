package main

import "fmt"

func init() {
    fmt.Println("dev add init()")
}

func main() {
    fmt.Println("main")
    
    fmt.Println("dev change main")
}

func foo() {
    fmt.Println("foo")
}

func bar() {
    fmt.Println("bar")
    fmt.Println("dev chang bar function")
}