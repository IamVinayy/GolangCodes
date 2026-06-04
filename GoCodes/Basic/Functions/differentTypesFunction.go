package main

import "fmt"

func main() {
	fmt.Println(add(1, 3)) // 1. called normal function.

	func() {
		fmt.Println("Hello, Anonymous Function.") // 2. anonymous function.
	}()

	greet := func() {
		fmt.Println("From greet, anonymous function")
	}

	greet()

	// 3. taking function as types:- assigned a function to a variable.

	operations := add

	fmt.Println(operations(3, 4))
}

// Normal Function - defined and called, if no return datatype specified, 0 is returned.

func add(a, b int) int {
	return a + b
}

/* 	anonymous function - function without name.
func (){
	fmt.Println("some text")
}()      [() - this is its call.]

*/

/*
	first class citizen - for functions - wide range of operations can be done on these.

*/
