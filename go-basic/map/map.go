package main

import "fmt"

// var ss []string = []string{}
func main() {
	var m = map[string]int{}
	fmt.Printf("the value: %#v\n", m)

	m["Answer"] = 42
	fmt.Printf("the value: %#v\n", m)

	v := m["Answer"]
	fmt.Printf("the value: %#v\n", v)

	m["Petch"] = 15
	fmt.Printf("the value: %#v\n", m)

	m["Petch"] = 0
	fmt.Printf("the value: %#v\n", m)

	// delete(m, "Petch")
	// fmt.Printf("the value: %#v\n", m)

	n, ok := m["Petch"]
	if ok {
		fmt.Printf("yes: %v\n", n)
	} else {
		fmt.Printf("no %v\n", n)
	}

	// var mapp map[string]int
	mapp := make(map[string]int)
	if mapp == nil{
		fmt.Println("m is nil")
	}
	fmt.Printf("The value: %#v\n", mapp)
}