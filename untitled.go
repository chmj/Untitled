package main

import "unsafe"

/*
   Some Facts in Go We Should Know

   Before introducing the valid unsafe pointer use patterns, we need to know some facts in Go.
   Fact 1: unsafe pointers are pointers and uintptr values are integers

   Each of non-nil safe and unsafe pointers references another value. However uintptr values don't reference any values,
   they are just plain integers, though often each of them stores an integer which can be used to represent a memory
   address.

   Go is a language supporting automatic garbage collection. When a Go program is running, Go runtime will check which
   memory blocks are not used by any value any more and collect the memory allocated for these unused blocks, from time
   to time. Pointers play an important role in the check process. If a memory block is unreachable from (referenced by)
   any values still in use, then Go runtime thinks it is an unused value and it can be safely garbage collected.

   As uintptr values are integers, they can participate arithmetic operations.

   The example in the next subsection shows the differences between pointers and uintptr values.
   Fact 2: unused memory blocks may be collected at any time

   At run time, the garbage collector may run at an uncertain time, and each garbage collection process may last an
   uncertain duration. So when a memory block becomes unused, it may be collected at an uncertain time.
*/

// Assume createInt will not be inlined.
//go:noinline
func createInt() *int {
	return new(int)
}

func main() {
	p0, y, z := createInt(), createInt(), createInt()
	var p1 = unsafe.Pointer(y)
	var p2 = uintptr(unsafe.Pointer(z))

	// At the time, even if the address of the int
	// value referenced by z is still stored in p2,
	// the int value has already become unused, so
	// garbage collector can collect the memory
	// allocated for it now. On the other hand, the
	// int values referenced by p0 and p1 are still
	// in use.

	// uintptr can participate arithmetic operations.
	p2 += 2
	p2--
	p2--

	*p0 = 1                         // okay
	*(*int)(p1) = 2                 // okay
	*(*int)(unsafe.Pointer(p2)) = 3 // dangerous!
}
