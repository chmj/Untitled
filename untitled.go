package main

import (
	"fmt"
)

/*
Struct Value Literals and Struct Value Manipulations

In Go, the form T{...}, where T must be a type literal or a type name, is called a composite literal and is used as the
value literals of some kinds of types, including struct types and the container types introduced later.

Note, a type literal T{...} is a typed value, its type is T.
Given a struct type S whose underlying type is struct{x int; y bool}, the zero value of S can be represented by the
following two variants of struct composite literal forms:

    S{0, false}. In this variant, no field names are present but all field values must be present by the field
declaration orders.
    S{x: 0, y: false}, S{y: false, x: 0}, S{x: 0}, S{y: false} and S{}. In this variant, each field item is optional
and the order of the field items is not important. The values of the absent fields will be set as the zero values of
their respective types. But if a field item is present, it must be presented with the FieldName: FieldValue form. The
order of the field items in this form doesn't matter. The form S{} is the most used zero value representation of type S.

If S is a struct type imported from another package, it is recommended to use the second form, to maintain
compatibility. Consider the case where the maintainer of the package adds a new field for type S, this will make the
use of first form invalid.

Surely, we can also use the struct composite literals to represent non-zero struct value.

For a value v of type S, we can use v.x and v.y, which are called selectors (or selector expressions), to represent the
field values of v. v is called the receiver of the selectors. Later, we call the dot . in a selector as the property
selection operator.
An example:
*/

type Book struct {
	title, author string
	pages         int
}

func main() {
	book := Book{"Go 101", "Tapir", 256}
	fmt.Println(book) // {Go 101 Tapir 256}

	// Create a book value with another form.
	// All of the three fields are specified.
	book = Book{author: "Tapir", pages: 256, title: "Go 101"}

	// None of the fields are specified. The title and
	// author fields are both "", pages field is 0.
	book = Book{}

	// Only specify the author field. The title field
	// is "" and the pages field is 0.
	book = Book{author: "Tapir"}

	// Initialize a struct value by using selectors.
	var book2 Book // <=> book2 := Book{}
	book2.author = "Tapir Liu"
	book2.pages = 300
	fmt.Println(book2.pages) // 300
}
