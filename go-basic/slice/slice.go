package main

import "fmt"

func show(tag string, sk []string) {
	l := len(sk)
	c := cap(sk)
	fmt.Printf("%s --- len: %d cap: %d --- show: %v\n", tag, l, c, sk)
}

func tips_trick() {
	x := [3]string{"js", "go", "python"}
	show("x", x[:])

	init := []string{"1", "2", "3"}
	show("init", init)
	a := init[0:2]
	a[1] = "4"
	show("init", a)

	a = append(a, "5")
	show("a", a)
	a[1] = "7"

	show("a", a)
	show("init", init)

	skills := []string{"js", "go", "python"}
	s1 := skills[0:2]
	show("s1", s1)

	s2 := skills[1:3]
	show("s2", s2)

	s2[0] = "gopherxx"
	show("s1", s1)
	show("s2", s2)

	fmt.Printf("skills: %#v\n", skills)
	// s3 := skillsCap[0:4]
	// show("s1", s3)
	s3 := skills[0:3]
	show("s3", s3)

	s4 := skills[1:3]
	s4 = append(s4, "c++")
	show("s4", s4)
	s4[0] = "gopher"
	show("s4", s4)
	// คำสั่ง append จะมีการสร้าง array ตัวใหม่ ให้ ก็ต่อเมื่อ capacity เกิดการขยายตัว ถ้าไม่ขยายจะยังอ้างอิง ตัวเดิม
}

func main() {
	// slice()
	tips_trick()
}

func slice() {
	skills := []string{"js", "go", "python"}
	for i := 0; i < len(skills); i++ {
		fmt.Printf("%#v\n", skills[i])
	}

	for _, v := range skills {
		fmt.Printf("for range value: %#v\n", v)
	}

	fmt.Printf("length: %d -- val: %#v\n", len(skills), skills)
	skills = append(skills, "ruby", "java", "erlang")

	show("skills", skills[0:2])
	show("skills", skills[1:])

	l := len(skills)
	show("skills", skills[0:l])
	show("skills", skills[0:])
	show("skills", skills[:l])
	show("skills", skills[:])

	var ss []int
	fmt.Printf("%#v\n", ss)
	if ss == nil {
		fmt.Println("nil")
	}

	v := make([]int, 3)
	fmt.Printf("%#v\n", v)
	v = append(v, 5)
	// v[10] = 36
	fmt.Printf("%#v\n", v)

	ss = append(ss, 33)
	fmt.Printf("%#v\n", ss)
}
