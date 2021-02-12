package main

import (
	"fmt"
	"time"
)

var week time.Duration

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
	t = time.Now().UTC()
	fmt.Println(t)
	fmt.Println(time.Now())
	// calculating times - giving some error
	// week = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
	// weekFronNow := t.Add(week)
	// fmt.Println(weekFronNow)
	// formatting times
	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format("02 Jan 2006 15:04"))
	s := t.Format("20060102")
	fmt.Println(t, "=>", s)
}

// 2020-12-27 15:22:07.299916655 +0530 IST m=+0.000063835
// 27.12.2020
// 2020-12-27 09:52:07.300157043 +0000 UTC
// 2020-12-27 15:22:07.300168744 +0530 IST m=+0.000315879
// 27 Dec 20 09:52 UTC
// Sun Dec 27 09:52:07 2020
// 27 Dec 2020 09:52
// 2020-12-27 09:52:07.300157043 +0000 UTC => 20201227
