package main

import "fmt"

func show(sk [3]string){
	fmt.Printf("show: %#v\n", sk)
}

func main() {
	skills := [3]string{"js", "go", "python"} 
	fmt.Printf("%#v\n", skills)
	s := skills[1]
	fmt.Printf("%#v\n", s)

	l := len(skills)
	fmt.Printf("len: %#v\n", l)

	skills[1] = "golang"
	fmt.Printf("%#v\n", skills)

	show(skills)

	var xs [3]string
	show(xs)
}