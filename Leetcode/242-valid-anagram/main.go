package main

import "fmt"

func isAnagram(s string, t string) bool {
	aph := [26]int{}
	for _, i := range s {
		aph[i-'a']++
	}
	for _, i := range t {
		aph[i-'a']--
	}

	return aph == [26]int{}
}

func main() {
	s, t := "nagaram", "anagram"
	fmt.Println(isAnagram(s, t))
	s, t = "rat", "car"
	fmt.Println(isAnagram(s, t))
}
