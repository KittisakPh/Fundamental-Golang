package json

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name       string `json:"name"`
	Level      string `json:"level"`
	Views      int    `json:"views"`
	Instructor string `json:"instructor"`
	FullPrice  int    `json:"full_price"`
}

func Json() {
	c := course{
		Name:       "basic go",
		Level:      "normal",
		Views:      7239,
		Instructor: "อนุชิโตะ",
		FullPrice:  9999,
	}

	// json marshal เป็นการ copy struct ไป เพื่อเข้าถึงใน Fields
	b, err := json.Marshal(c)
	fmt.Printf("type: %T\n", b)
	fmt.Printf("byte: %v\n", string(b))
	fmt.Printf("string: %s\n", b)
	if err != nil {
		fmt.Printf("%#v\n", c)
	}

	// json unmarshal เป็นการ เพิ่มค่าใน struct เลยต้องส่ง address ไป
	data := []byte(`{
		"Name": "basic go",
		"Level": "normal",
		"Views": 7239,
		"Instructor": "อนุชิโตะ",
		"FullPrice": 9999
	}`)
	var cStruct course
	err = json.Unmarshal(data, &cStruct)
	if err != nil {
		fmt.Printf("%#v\n", c)
	}
	fmt.Println(cStruct)
}