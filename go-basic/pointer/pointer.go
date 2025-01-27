package main

import "fmt"

type course struct {
	name, instructor string
	price            int
}

func (c course)discount() int {
	c.price = c.price - 599
	fmt.Println("discount:", c.price)
	return c.price
}

func (c *course)discountPtr() int {
	// (*c).price = (*c).price - 599
	c.price = c.price - 599
	fmt.Println("discount:", c.price)
	return c.price
}

func main() {
	c := new(course)
	fmt.Printf("%#v\n", c)
	c.price = 5000
	// c := &course{"Basic Go", "AnuchitO", 9999}
	d := c.discount()
	fmt.Println("discount price:", d)
	fmt.Println("price:", c.price)

	e := c.discountPtr()
	fmt.Println("discountPtr price:", e)
	fmt.Println("price:", c.price)

	// nil => zero value of pointer
	var addr *int
	fmt.Printf("zero value: %v", addr)
}

func changePrice(p int) {
	p -= 599
	fmt.Println("change: ", p, &p)
}

func changePricePtr(p *int) {
	*p -= 599
	fmt.Println("change: ", *p, p)
}

func pointer() {
	var price int = 9999
	var addr *int = &price
	*addr = 9400 // write
	fmt.Println(price, addr)
	fmt.Printf("type: %T\n", addr)
	v := *addr //read
	fmt.Println(v)

	changePrice(price)
	fmt.Println("after: ", price, addr)

	changePricePtr(&price)
	fmt.Println("after: ", price, addr)
}
