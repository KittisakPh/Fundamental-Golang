package constant

import "fmt"

type day int

const Pi = true

const (
	_ = iota
	Sunday day = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func Constant() {
	// var world string = "เพชร"
	// fmt.Println("Hello", world)
	// fmt.Println("Happy", Pi, "Day")

	// world = "Petch"
	// fmt.Println("Hello", world)

	fmt.Println("Sunday :", Sunday)
	fmt.Println("Monday :", Monday)
	fmt.Println("Tuesday :", Tuesday)
	fmt.Println("Wednesday :", Wednesday)
	fmt.Println("Thursday :", Thursday)
	fmt.Println("Fiday :", Friday)
	fmt.Println("Saturday :", Saturday)
	Skills("js", "python", "go")
}

func Skills(xs ...string){
	fmt.Println(len(xs))
	fmt.Println(xs)
	// fmt.Println(xs[0])
	// fmt.Println(xs[1])
	// fmt.Println(xs[2])
	if xs == nil{
		fmt.Println("nil")
	}

	for _, s := range xs{
		fmt.Println("skills:", s)
	}
}