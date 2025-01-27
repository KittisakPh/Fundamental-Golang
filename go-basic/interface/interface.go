package main

import "fmt"

// ข้อตกลง
type xxx interface {
	discount() int
	info()
}

func summary(val xxx) {
	fmt.Println("discount price:", val.discount())
	val.info()
}

type promotioner interface {
	discount() int
}

type infoer interface {
	info()
}

func sale(val promotioner) {
	fmt.Printf("sale: %#v\n", val.discount())
}

func t(val infoer) {
	val.info()
}

type infoerT struct{}

func (a infoerT) info() {
	fmt.Println("1234")
}

func show(val any) {
	// i, ok := val.(int)
	// if ok {
	// 	i += 2
	// 	fmt.Println(i)
	// } else {
	// 	fmt.Println("not int")
	// }

	// j, k := val.(string)
	// if k {
	// 	j += " hello"
	// 	fmt.Println(j)
	// }

	switch v := val.(type) {
	case int:
		i := v + 2
		fmt.Println("int:", i)
	case string:
		s := v + " hello"
		fmt.Println("string:", s)
	default:
		fmt.Println("not handle type.")
	}
}

type course struct{}

func (c course) discount() int {
	return 599
}

func (c course) info() {
	fmt.Println("info:", c)
}

func main() {
	var v course
	// v = 36
	sale(v)

	summary(v)

	// fmt.Printf("%T %#v\n", v, v)
	// fmt.Printf("%T %#v\n", v, v)
	// v = "gopher"
	// show(&v)

	// x := "gopher"
	// sale(x)
	var info infoerT
	t(info)
}
