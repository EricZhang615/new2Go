package main

import (
	"fmt"
	"time"
)

func main() {
	// var
	//var a = "asdasdad"
	//
	//var b, c int = 1, 2
	//
	//var d = true
	//
	//var e float64
	//
	//f := float32(e)
	//
	//g := a + "foo"
	//
	//fmt.Println(a, b, c, d, e, f)
	//fmt.Println(g)
	//
	//const s string = "const"
	//
	//const h = 500000000000
	//
	////const i = 3e20 / h
	//
	////fmt.Println(s, h, i, math.Sin(h), math.Sin(i))
	//
	//fmt.Println("hello world")

	// for
	i := 1
	for {
		fmt.Println("loop")
		break
	}
	for j := 7; j < 9; j++ {
		fmt.Println(j)
	}
	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// if
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is neg")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// switch
	a := 2
	switch a {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4, 5:
		fmt.Println("four or five")
	default:
		fmt.Println("other")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("before noon")
	default:
		fmt.Println("afternoon")
	}

	// array
	var aArr [5]int
	aArr[4] = 100
	fmt.Println("get:", aArr[2])
	fmt.Println("len:", len(aArr))
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
