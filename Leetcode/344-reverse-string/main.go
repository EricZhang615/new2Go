package main

import "fmt"

func reverseString(s []byte) {
	i, j := 0, len(s)-1
	for i < j {
		tmp := s[i]
		s[i] = s[j]
		s[j] = tmp
		i++
		j--
	}
}

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	fmt.Println(s)
}
