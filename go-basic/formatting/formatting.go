package main

import "fmt"

func main(){
	msg, age, price, check := "Hello Gopher!!!", 55, 22.52, true
	var r rune = 'A'

	fmt.Printf("Type: %T -- msg: %#v\n", msg, msg)
	fmt.Printf("Type: %T -- age: %#v\n", age, age)
	fmt.Printf("Type: %T -- price: %#v\n", price, price)
	fmt.Printf("Type: %T -- check: %#v\n", check, check)
	fmt.Printf("Type: %T -- r: %c\n", r, r)
}