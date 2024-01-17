package main

import "fmt"

// find-maximum-number-of-string-pairs

func maximumNumberOfStringPairs(words []string) int {
	set := make(map[string]int)
	res := 0
	for i := 0; i < len(words); i++ {
		set[words[i]] = i + 1
	}
	for i := 0; i < len(words); i++ {
		str := string(words[i][1]) + string(words[i][0])
		if set[str] != 0 && set[str] <= i {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(maximumNumberOfStringPairs([]string{"cd", "ac", "dc", "ca", "zz"}))
}
