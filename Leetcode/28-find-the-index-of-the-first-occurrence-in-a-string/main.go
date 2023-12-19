package main

import "fmt"

func strStr(haystack string, needle string) int {
	next := make([]int, len(needle))
	j := -1
	next[0] = j
	for i := 1; i < len(needle); i++ {
		for j >= 0 && needle[j+1] != needle[i] {
			j = next[j]
		}
		if needle[j+1] == needle[i] {
			j++
		}
		next[i] = j
	}
	j = -1
	for i := 0; i < len(haystack); i++ {
		for j >= 0 && haystack[i] != needle[j+1] {
			j = next[j]
		}
		if haystack[i] == needle[j+1] {
			j++
		}
		if j == len(needle)-1 {
			return i - len(needle) + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("leetcode", "leeto"))
}
