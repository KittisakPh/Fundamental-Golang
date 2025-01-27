package main

import "fmt"

// Workshop: Println
// Output:
// เรื่อง: Avengers: Endgame
// ปี: 2019
// เรตติ้ง: 8.4
// ประเภท: Sci-Fi
// ซุปเปอร์ฮีโร่: true
// TODO: write code below.

func main() {
	// ->use variable
	title, year, ratings, genre, isSuperHero := "Avengers: Endgame", 2019, 8.4, "Sci-Fi", true
	fmt.Println("เรื่อง:", title)
	fmt.Println("ปี:", year)
	fmt.Println("เรตติ้ง:", ratings)
	fmt.Println("ประเภท:", genre)
	fmt.Println("ซุปเปอร์ฮีโร่:", isSuperHero)

	// ->use Raw string
	movie :=`เรื่อง: Avengers: Endgame
ปี: 2019
เรตติ้ง: 8.4
ประเภท: Sci-Fi
ซุปเปอร์ฮีโร่: true`
	fmt.Println(movie)
}