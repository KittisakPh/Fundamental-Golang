package main

import (
	"fmt"
	"math"
)

func main() {
	// num := 33
	// if num == 34 && (num >= 30 || num < 39) {
	// 	fmt.Println("Yes!! it's Thirty four	")
	// }
	// if num%2 == 0 {
	// 	fmt.Println("Yes!! it's even.")
	// } else if num == 35 {
	// 	fmt.Println("it's Thirty five.")
	// } else {
	// 	fmt.Println("No!! it's odd.")
	// }

	lim := 225.0
	if v := math.Pow(10, 2); v < lim {
		fmt.Println("x power n is:", v)
	} else {
		fmt.Printf("x power n is %g over limit %g\n", v, lim)
	}
	// fmt.Println(v)
}