package main

import "fmt"

func init() {
	fmt.Println("dev add init()")
}

func main() {
	fmt.Println("main")
	fmt.Println("master change main")
	fmt.Println("test branch change main")

}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
