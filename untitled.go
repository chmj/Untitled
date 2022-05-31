package main

import "fmt"

/*
Pointer Dereference

Given a pointer value p of a pointer type whose base type is T, how can you get the value at the address stored in the
pointer (a.k.a., the value being referenced by the pointer)? Just use the expression *p, where * is called dereference
operator. *p is called the dereference of pointer p. Pointer dereference is the inverse process of address taking.
The result of *p is a value of type T (the base type of the type of p).

Dereferencing a nil pointer causes a runtime panic.
The following program shows some address taking and pointer dereference examples:
*/

func main() {
	p0 := new(int)   // p0 points to a zero int value.
	fmt.Println(p0)  // (a hex address string)
	fmt.Println(*p0) // 0

	// x is a copy of the value at
	// the address stored in p0.
	x := *p0
	// Both take the address of x.
	// x, *p1 and *p2 represent the same value.
	p1, p2 := &x, &x
	fmt.Println(p1 == p2) // true
	fmt.Println(p0 == p1) // false
	p3 := &*p0            // <=> p3 := &(*p0) <=> p3 := p0
	// Now, p3 and p0 store the same address.
	fmt.Println(p0 == p3) // true
	*p0, *p1 = 123, 789
	fmt.Println(*p2, x, *p3) // 789 789 123

	fmt.Printf("%T, %T \n", *p0, x) // int, int
	fmt.Printf("%T, %T \n", p0, p1) // *int, *int
}
