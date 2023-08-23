package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// func
func add(a int, b int) int {
	return a + b
}

//func add2(a, b int) int {
//	return a + b
//}

func exists(m map[string]string, k string) (v string, ok bool) {
	v, ok = m[k]
	return v, ok
}

// point
func add2(n int) {
	n += 2
}

func add2ptr(n *int) {
	*n += 2
}

// struct
type user struct {
	name     string
	password string
}

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

	// slice
	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)

	fmt.Println(s[2:5])
	fmt.Println(s[:5])
	fmt.Println(s[2:])

	good := []string{"g", "o", "o", "d"}
	fmt.Println(good)

	// map
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(m["one"])
	fmt.Println(m["unknown"])

	r, ok := m["unknown"]
	fmt.Println(r, ok)

	delete(m, "one")

	m2 := map[string]int{"one": 1, "two": 2}
	var m3 = map[string]int{"one": 1, "two": 2}
	fmt.Println(m2, m3)

	// range
	nums := []int{2, 3, 4}
	sum := 0
	for i, num := range nums {
		sum += num
		if num == 2 {
			fmt.Println("index:", i, "num:", num)
		}
	}
	fmt.Println(sum)

	m4 := map[string]string{"a": "A", "b": "B"} // map is disordered
	for k, v := range m4 {
		fmt.Println(k, v)
	}
	for k := range m4 {
		fmt.Println("key", k)
	}

	// func
	res := add(1, 2)
	fmt.Println(res)

	v, ok := exists(map[string]string{"a": "A"}, "a")
	fmt.Println(v, ok)

	// point
	n := 5
	add2(n)
	fmt.Println(n)
	add2ptr(&n)
	fmt.Println(n)

	// struct
	sa := user{name: "wang", password: "1024"}
	sb := user{"wang", "1024"}
	sc := user{name: "wang"}
	sc.password = "1024"
	var sd user
	sd.name = "wang"
	sd.password = "1024"

	fmt.Println(sa, sb, sc, sd)
	fmt.Println(checkPassword(sa, "haha"))
	fmt.Println(checkPassword2(&sa, "haha"))

	// struct method
	sa.resetPassword("2048")
	fmt.Println(sa.checkPassword("2048"))
	// 结构体是值类型
	// (*u).password 可以简写为 u.password

	// error
	u, err := findUser([]user{{"wang", "1024"}}, "wang")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.name)

	if u, err := findUser([]user{{"wang", "1024"}}, "li"); err != nil {
		fmt.Println(err)
		// return
	} else {
		fmt.Println(u.name)
	}

	// string
	aa := "hello"
	fmt.Println(strings.Contains(aa, "ll"))               // true
	fmt.Println(strings.Count(aa, "l"))                   // 2
	fmt.Println(strings.HasPrefix(aa, "he"))              // true
	fmt.Println(strings.HasSuffix(aa, "llo"))             // true
	fmt.Println(strings.Index(aa, "ll"))                  // 2
	fmt.Println(strings.Join([]string{"he", "llo"}, "-")) // he-llo
	fmt.Println(strings.Repeat(aa, 2))                    // hellohello
	fmt.Println(strings.Replace(aa, "e", "E", -1))        // hEllo
	fmt.Println(strings.Split("a-b-c", "-"))              // [a b c]
	fmt.Println(strings.ToLower(aa))                      // hello
	fmt.Println(strings.ToUpper(aa))                      // HELLO
	fmt.Println(len(aa))                                  // 5
	bb := "你好"
	fmt.Println(len(bb)) // 6
}

// struct
func checkPassword(u user, password string) bool {
	return u.password == password
}

func checkPassword2(u *user, password string) bool {
	return u.password == password
}

// struct method
func (u user) checkPassword(password string) bool {
	return u.password == password
}

func (u *user) resetPassword(password string) {
	u.password = password
}

// error
func findUser(users []user, name string) (v *user, err error) {
	for _, u := range users {
		if u.name == name {
			return &u, nil
		}
	}
	return nil, errors.New("not found")
}
