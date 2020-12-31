package main

import "fmt"

func init() {
	fmt.Println("dev add init()")
}

func main() {
	fmt.Println("main")
	fmt.Println("master change main")
	fmt.Println("test branch change main")

	fmt.Println("master chage main again")
	fmt.Println("lef foo")
	fmt.Println("master foo")

}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
	fmt.Println("lef foo")
}
