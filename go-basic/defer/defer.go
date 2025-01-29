package deferr

import "fmt"

func Defer() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
	fmt.Println("done")
}