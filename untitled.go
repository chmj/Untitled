package main

import "fmt"

/*
Panic and Recover

Go doesn't support exception throwing and catching, instead explicit error handling is preferred to use in Go
programming. In fact, Go supports an exception throw/catch alike mechanism. The mechanism is called panic/recover.

We can call the built-in panic function to create a panic to make the current goroutine enter panicking status.

Panicking is another way to make a function return. Once a panic occurs in a function call, the function call returns
immediately and enters its exiting phase.

By calling the built-in recover function in a deferred call, an alive panic in the current goroutine can be removed so
that the current goroutine will enter normal calm status again.

If a panicking goroutine exits without being recovered, it will make the whole program crash.
The built-in panic and recover functions are declared as:

func panic(v interface{})
func recover() interface{}

Interface types and values will be explained in the article interfaces in Go later. Here, we just need to know that the
blank interface type interface{} can be viewed as the "Any" type or the Object type in many other languages. In other
words, we can pass a value of any type to a panic function call.

The value returned by a recover function call is the value a panic function call consumed.
The example below shows how to create a panic and how to recover from it.
*/

func main() {
	defer func() {
		fmt.Println("exit normally.")
	}()
	fmt.Println("hi!")
	defer func() {
		v := recover()
		fmt.Println("recovered:", v)
	}()
	panic("bye!")
	fmt.Println("unreachable")
}
