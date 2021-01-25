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
	fmt.Println("matser foo")
}

func bar() {
	fmt.Println("bar")
	fmt.Println("lef foo")
}

func baz() {
	fmt.Println("baz")
	fmt.Println("baz again")
	fmt.Println("baz 3rd")
	fmt.Println("baz 4th")
	fmt.Println("baz 5th")
	fmt.Println("lo change baz")
}

func lo() {
	fmt.Println("content we want to merge")
	fmt.Println("lo change lo")

}
