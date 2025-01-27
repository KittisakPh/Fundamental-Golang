package main

import "fmt"

// custom type
type course struct {
	name       string
	instructor string
	price      float64
}

func main() {
	c := course{name: "Basic Go", instructor: "AnuchitO", price: 9999}
	var c2 = course{"Basic Go", "Petch", 9}
	var c3 = course{instructor: "kts"}
	var c4 = course{}

	fmt.Printf("%v\n", c2)
	fmt.Printf("%#v\n", c3)
	fmt.Printf("%#v\n", c4)

	n := c.name
	c.instructor = "Petch"

	fmt.Printf("name: %s\n", n)
	fmt.Printf("instructor: %s\n", c.instructor)
	fmt.Printf("price: %.2f\n", c.price)
	fmt.Printf("course: %v\n", c)
}
