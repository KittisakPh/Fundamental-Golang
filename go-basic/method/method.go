package main

import "fmt"

type course struct {
	name, instructor string
	price            int
}

func (c course) discount(d int) int {
	p := c.price - d
	return p
}

func (c course) info() {
	fmt.Println("name:", c.name)
	fmt.Println("instructor:", c.instructor)
	fmt.Println("price:", c.price)
}

func main() {
	c := course{"Basic Go", "AnuchitO", 9999}
	fmt.Printf("%#v\n", c)

	d := c.discount(599)
	fmt.Println("discount price:", d)

	c.info()
}