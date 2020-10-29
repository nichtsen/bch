package main

import "fmt"

func init() {
	fmt.Println("dev add init()")
}

func main() {
	fmt.Println("main")
<<<<<<< HEAD
	fmt.Pritnln("main branch change main func")

	fmt.Println("dev change main")
=======

	fmt.Println("my branch change main")
>>>>>>> my
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
	fmt.Println("dev chang bar function")
}
