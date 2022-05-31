package main

import "fmt"

func main() {
	// Nested Map[struct]->Map[string]->int
	languages := map[struct{ dynamic, strong bool }]map[string]int{
		{true, false}:  {"JavaScript": 1995},
		{false, true}:  {"Go": 2009},
		{false, false}: {"C": 1972},
	}
	// The key type and element type of this map
	// are both pointer types. Some weird, just
	// for education purpose.
	m0 := map[*struct{ dynamic, strong bool }]*map[string]int{}
	for category, langInfo := range languages {
		m0[&category] = &langInfo
		// This following line has no effects on languages.
		category.dynamic, category.strong = true, true
	}
	for category, langInfo := range languages {
		fmt.Println(category, langInfo)
	}

	m1 := map[struct{ dynamic, strong bool }]map[string]int{}
	for category, langInfo := range m0 {
		m1[*category] = *langInfo
	}
	// m0 and m1 both contain only one entry.
	fmt.Println(len(m0), len(m1)) // 1 1
	fmt.Println(m1)               // map[{true true}:map[C:1972]]
}
