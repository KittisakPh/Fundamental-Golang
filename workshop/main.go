package main

import "fmt"

func main() {
	// variable()
	// formatting()
	// if_else()
	// switch_case()
	// function()
	// array()
	// for_loop()
	slice()
}

// ===========================================================
// Workshop: slice
// กำหนด: 1. เราได้เก็บสะสมคะแนนโหวตผู้ชมไว้เป็น 2 ชุด ที่เก็บอยู่ในตัวแปร xs และ ys เรียบร้อยแล้ว
// 	  2. ให้ทำการรวมคะแนนโหวตที่อยู่ในตัวแปร xs และ ys ไปเก็บไว้ในตัวแปร votes ตามลำดับ (ค่าใน xs ทั้งหมดแล้วต่อด้วย ys).
//	  3. ทำการแสดงผลคะแนนโหวตของผู้ชมที่อยู่ในตำแหน่ง(index)ที่ 5 ถึง 8 ของ votes ออกมาทางหน้าจอ
//
// Output:
// votes: [4 5 7 8 3 8 0 7 2 10 9 7]
// votes 5 - 8: [8 0 7 2]
// ===========================================================
func slice() {
	xs := []float64{4, 5, 7, 8, 3, 8, 0}
	ys := []float64{7, 2, 10, 9, 7}

	var votes []float64
	votes = append(votes, xs...)
	votes = append(votes, ys...)

	fmt.Printf("votes: %v\n", votes)
	fmt.Printf("votes 5 - 8: %v\n", votes[5:9])
}

// ===========================================================
// Workshop: for-loop
// 1. ให้ใช้ for loop ทำการเปลี่ยนค่าในอาร์เรย์ genres โดยเพิ่มคำนำหน้า "Movie: " ทุกค่า ดังนี้ "Movie: Action", "Movie: Adventure", “Movie: Fantasy”
// 2. ให้แสดงผลค่า genres ทีละค่า โดยใช้ for-range
// Output:
// before for loop: [3]string{"Action", "Adventure", "Fantasy"}
// after for loop: [3]string{"Movie: Action", "Movie: Adventure", "Movie: Fantasy"}
// genres[0]: Movie: Action
// genres[1]: Movie: Adventure
// genres[2]: Movie: Fantasy
// ===========================================================
func for_loop() {
	genres := [3]string{"Action", "Adventure", "Fantasy"}
	fmt.Printf("before for loop: %#v\n", genres)
	for i := 0; i < len(genres); i++ {
		genres[i] = "Movie: " + genres[i]
	}
	fmt.Printf("after for loop: %#v\n", genres)

	for i, v := range genres {
		fmt.Printf("genres[%d]: %s\n", i, v)
	}
}

// ===========================================================
// Workshop: array
// ให้ประกาศตัวแปรอาร์เรย์ประเภทหนัง(genres)ที่เก็บค่าเป็น "Action", “Adventure” และ “Fantasy” ตามลำดับ
// ให้อ่านค่าของอาร์เรย์แต่ละช่องเพื่อแสดงผล ให้แสดงผลทั้งหมดตามตัวอย่าง Output ด้านล่าง
// หลังจากนั้นเปลี่ยนแปลงค่าในอาร์เรย์ index ที่ 1 ให้เป็น Sci-Fi พร้อมทั้งแสดงผล เพื่อยืนยันว่าค่าเปลี่ยนจริง
// Output:
// “genres[0]: Action”
// “genres[1]: Adventure”
// “genres[2]: Fantasy”
// “genres[1]: Sci-Fi”
// ===========================================================
func array() {
	// normal
	{
		genres := [3]string{"Action", "Adventure", "Fantasy"}
		fmt.Printf("genres[0]: %s\n", genres[0])
		fmt.Printf("genres[1]: %s\n", genres[1])
		fmt.Printf("genres[2]: %s\n", genres[2])
		genres[1] = "Sci-Fi"
		fmt.Printf("genres[1]: %s\n", genres[1])
	}
	// variadic
	{
		genres := [...]string{"Action", "Adventure", "Fantasy"}
		fmt.Printf("genres[0]: %s\n", genres[0])
		fmt.Printf("genres[1]: %s\n", genres[1])
		fmt.Printf("genres[2]: %s\n", genres[2])
		genres[1] = "Sci-Fi"
		fmt.Printf("genres[1]: %s\n", genres[1])
	}
}

// ===========================================================
// Workshop: function
// สร้างฟังก์ชันชื่อ emote และรับพารามิเตอร์หนึ่งตัวชื่อ ratings เป็นตัวแปรชนิด float64 และคืนค่าเป็น string
// ถ้า ratings มีค่าน้อยกว่า 5.0 จะทำการคืนค่าคำว่า “Disappointed 😞”
// ถ้า ratings มีค่ามากกว่าหรือเท่ากับ 5.0 และน้อยกว่า 7.0 จะทำการคืนค่าคำว่า “Normal 😐”
// ถ้า ratings มีค่ามากกว่าหรือเท่ากับ 7.0 และน้อยกว่า 10.0 จะทำการคืนค่าคำว่า “Good 🥰”
// กรณีอื่นๆ ให้ทำการคืนค่าคำว่า “🤷🤷🤷🤷”
// ===========================================================
func emote(ratings float64) string {
	switch {
	case ratings < 5.0:
		return "Disappointed 😞"
	case ratings >= 5.0 && ratings < 7.0:
		return "Normal 😐"
	case ratings >= 7.0 && ratings < 10.0:
		return "Good 🥰"
	}
	return "🤷🤷🤷🤷"
}
func function() {
	fmt.Println(emote(4.9))
	fmt.Println(emote(5.0))
	fmt.Println(emote(7.0))
	fmt.Println(emote(15.0))
}

// ===========================================================
// Workshop: switch-case
// ratings := 8.4
// ถ้า ratings มีค่าน้อยกว่า 5.0 จะทำการแสดงผลคำว่า “Disappointed 😞”
// ถ้า ratings มีค่ามากกว่าหรือเท่ากับ 5.0 และน้อยกว่า 7.0 จะทำการแสดงผลคำว่า “Normal 😐”
// ถ้า ratings มีค่ามากกว่าหรือเท่ากับ 7.0 และน้อยกว่า 10.0 จะทำการแสดงผลคำว่า “Good 🥰”
// กรณีอื่นๆ ให้ทำการแสดงผลคำว่า “🤷🤷🤷🤷”
// ===========================================================
func switch_case() {
	ratings := 8.4
	switch {
	case ratings < 5.0:
		fmt.Println("Disappointed 😞")
	case ratings >= 5.0 && ratings < 7.0:
		fmt.Println("Normal 😐")
	case ratings >= 7.0 && ratings < 10.0:
		fmt.Println("Good 🥰")
	default:
		fmt.Println("🤷🤷🤷🤷")
	}
}

// ===========================================================
// Workshop: if-else
// ถ้า ratings มีค่าน้อยกว่า 5.0 จะทำการแสดงผลคำว่า “Disappointed 😞”
// ถ้า ratings มีค่ามากกว่าหรือเท่ากับ 5.0 และน้อยกว่า 7.0 จะทำการแสดงผลคำว่า “Normal 😐”
// ถ้า ratings มีค่ามากกว่าหรือเท่ากับ 7.0 และน้อยกว่า 10.0 จะทำการแสดงผลคำว่า “Good 🥰”
// กรณีอื่นๆ ให้ทำการแสดงผลคำว่า “🤷🤷🤷🤷”
// ===========================================================
func if_else() {
	if ratings := 10; ratings < 5.0 {
		fmt.Printf("Disappointed %c\n", '😞')
	} else if ratings >= 5.0 && ratings < 7.0 {
		fmt.Printf("Normal %c\n", '😐')
	} else if ratings >= 7.0 && ratings < 10.0 {
		fmt.Printf("Good %c\n", '🥰')
	} else {
		fmt.Printf("%s\n", "🤷🤷🤷🤷")
	}
}

// ===========================================================
// Workshop: Formatting (Printf)
// เรื่อง: Avengers: Endgame
// ปี: 2019
// เรตติ้ง: 8.4
// ประเภท: Sci-Fi
// ซุปเปอร์ฮีโร่: true
// เรื่องโปรด: ⭐
// ===========================================================
func formatting() {
	// row string
	{
		movie := `เรื่อง: Avengers: Endgame
ปี: 2019
เรตติ้ง: 8.4
ประเภท: Sci-Fi
ซุปเปอร์ฮีโร่: true
เรื่องโปรด: ⭐
`
		fmt.Printf("%s\n", movie)
	}
	// literal
	{
		title, year, ratings, genre, isSuperHero, favorite := "Avengers: Endgame", 2019, 8.4, "Sci-Fi", true, '⭐'
		fmt.Printf("เรื่อง: %s\n", title)
		fmt.Printf("ปี: %d\n", year)
		fmt.Printf("เรตติ้ง: %.1f\n", ratings)
		fmt.Printf("ประเภท %s\n", genre)
		fmt.Printf("ซุปเปอร์ฮีโร่: %t\n", isSuperHero)
		fmt.Printf("เรื่องโปรด: %c\n", favorite)
	}
}

// ===========================================================
// Workshop: Variable (Println)
// เรื่อง: Avengers: Endgame
// ปี: 2019
// เรตติ้ง: 8.4
// ประเภท: Sci-Fi
// ซุปเปอร์ฮีโร่: true
// ===========================================================
func variable() {
	// variable
	{
		var title, year, ratings, genre, isSuperHero, favorite = "Avengers: Endgame", 2019, 8.4, "Sci-Fi", true, '⭐'
		fmt.Printf("เรื่อง: %v\n", title)
		fmt.Printf("ปี: %d\n", year)
		fmt.Printf("เรตติ้ง: %v\n", ratings)
		fmt.Printf("ประเภท: %v\n", genre)
		fmt.Printf("ซุปเปอร์ฮีโร่: %v\n", isSuperHero)
		fmt.Printf("เรื่องโปรด: %c\n", favorite)
	}
	// literal
	{
		title, year, ratings, genre, isSuperHero := "Avengers: Endgame", 2019, 8.4, "Sci-Fi", true
		fmt.Println("เรื่อง:", title)
		fmt.Println("ปี:", year)
		fmt.Println("เรตติ้ง:", ratings)
		fmt.Println("ประเภท:", genre)
		fmt.Println("ซุปเปอร์ฮีโร่:", isSuperHero)
	}
	// Raw string
	{
		movie := `เรื่อง: Avengers: Endgame
		ปี: 2019
		เรตติ้ง: 8.4
		ประเภท: Sci-Fi
		ซุปเปอร์ฮีโร่: true`
		fmt.Println(movie)
	}
}