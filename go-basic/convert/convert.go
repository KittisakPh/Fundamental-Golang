package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 256
	fmt.Printf("type:%T, val: %d\n", i, i)

	var f float64 = float64(i)
	fmt.Printf("type:%T, val: %f\n", f, f)

	var u uint8 = uint8(f)
	fmt.Printf("type:%T, val: %d\n", u, u)

	v := "72"
	i, err := strconv.Atoi(v)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(i)

	k := 10
	s := strconv.Itoa(k)
	fmt.Println(s)
}