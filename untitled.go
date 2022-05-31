package main

import "fmt"

/*
Variadic function calls
There are two manners to pass arguments to a variadic parameter of type []T:

    pass a slice value as the only argument. The slice must be assignable to values of type []T, and the slice must be
	followed by three dots .... The passed slice is called as a variadic argument.
    pass zero or more arguments which are assignable to values of type T. These arguments will be copied (or converted)
	as the elements of a new allocated slice value of type []T, then the new allocated slice will be passed to the
	variadic parameter.

Note, the two manners can't be mixed in the same variadic function call.
An example program which uses some variadic function calls:
*/

func Concat(sep string, tokens ...string) (r string) {
	for i, t := range tokens {
		if i != 0 {
			r += sep
		}
		r += t
	}
	return
}

func main() {
	tokens := []string{"Go", "C", "Rust"}
	// manner 1
	langsA := Concat(",", tokens...)
	// manner 2
	langsB := Concat(",", "Go", "C", "Rust")
	fmt.Println(langsA == langsB) // true
}
