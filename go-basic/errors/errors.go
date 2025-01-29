package main

import (
	"errors"
	"fmt"
)

// zero value interface = nil

type MyError struct{}

var myErr = errors.New("my custom error message")

func (e MyError) Error() string {
	return "MyError"
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, myErr
	}
	r := a / b
	return r, nil
}

func main() {
	r, err := Divide(1, 0)
	if err != nil {
		fmt.Println("handler err:", err)
		return
	}
	fmt.Println(r, err)
}
