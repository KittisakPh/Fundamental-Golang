package main

import "fmt"

func main() {
	switch today := "Tuesday"; today {
	case "Saturday":
		fmt.Println("today is Saturday")
		// fallthrough
	case "Monday", "Tuesday":
		fmt.Println("today is weekdays")
		// fallthrough
	default:
		fmt.Printf("สวัสดีวัน %v อยู่ดีมีแฮงเดย์\n", today)
	}

	num := 1
	switch{
	case num < 0:
		fmt.Println("negative number.")
	case num == 0:
		fmt.Println("zero number.")
	case num > 0:
		fmt.Println("positive number.")
	default:
		fmt.Println("🤷🤷🤷🤷")
	}	
}