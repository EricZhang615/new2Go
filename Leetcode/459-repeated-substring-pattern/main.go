package main

import "fmt"

func repeatedSubstringPattern(s string) bool {
	next := make([]int, len(s))
	j := -1
	next[0] = j
	for i := 1; i < len(s); i++ {
		for j >= 0 && s[j+1] != s[i] {
			j = next[j]
		}
		if s[j+1] == s[i] {
			j++
		}
		next[i] = j
	}
	if next[len(s)-1] != -1 && len(s)%(len(s)-(next[len(s)-1]+1)) == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(repeatedSubstringPattern("abac"))
}
