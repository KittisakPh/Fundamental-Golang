package main

import "fmt"

func main() {
	sum := 1
	// white-loop
	for sum < 5 {
		sum += sum
		fmt.Println("sum:", sum)
	}
	fmt.Println("sum done:", sum)

	// for {
	// 	fmt.Println("print forever.")
	// }

	skills := [3]string{"js", "go", "python"}
	for i := 0; i < len(skills); i++ {
		fmt.Println(skills[i])
	}

	// for-range
	for _, v := range skills {
		fmt.Println("value:", v)
		// fmt.Println("index:", i, "value:", v)
	}
}