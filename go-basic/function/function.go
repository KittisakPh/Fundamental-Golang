package main

import (
	"fmt"
	"math"
)

var x = "Hello, World!"
var y = 4

// var add = func(x, y float64) float64 { return x + y }
func add(x, y float64) float64 { return x + y }

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func compute(fn func(float64, float64) float64) float64 {
	v := fn(42, 13)
	return v
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x int, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	// a, b := add(42, 13)
	// fmt.Println(a, b)

	// a,b := swap("hello", "word")
	// fmt.Println(a, b)

	// v := add(42, 13)
	// fmt.Println(v)
	// fmt.Printf("%T\n", add)

	// r := compute(add)
	// fmt.Println("r:", r)

	// x := compute(hypot)
	// fmt.Println("x:", x)

	inc, curr := adder()
	x := inc()
	fmt.Println(x)
	x = inc()
	y := curr()
	fmt.Println(y)
}

// higher-order function
func adder() (func() int, func() int) {
	sum := 0
	return func() int {
		sum = sum + 1
			return sum
		}, func() int {
			return sum
		}
}
