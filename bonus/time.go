package bonus

import (
	"fmt"
	"time"
)

type day int

func Time() {
	now := time.Now()
	fmt.Println(now)
	// create time
	t := time.Date(2022, 07, 30, 20, 40, 58, 0, time.UTC)
	format := t.Format("02/01/2006 15:04:05")
	fmt.Println(format) //Output: 30/07/2022 20:40:58

	// convert string to time
	layout := "2006-01-02T15:04:05"
	t, _ = time.Parse(layout, format)
	fmt.Println(t)
}