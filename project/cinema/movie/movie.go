package movie

import (
	"fmt"
	"github.com/KittisakPh/cinema/ticket"
)

func init() {
	fmt.Println("init: movie")
}

func Review(name string, rating float64) {
	fmt.Printf("!!! I reviewed %s, it's rating %f\n", name, rating)
	ticket.BuyTicket(name)
}
