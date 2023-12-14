package main

import "fmt"

func findAnagrams(s string, p string) []int {
	aph := [26]int{}
	var ans []int
	for _, i := range p {
		aph[i-'a']++
	}
	i, j := 0, len(p)
	for j <= len(s) {
		tmp := [26]int{}
		for _, k := range s[i:j] {
			tmp[k-'a']++
		}
		if tmp == aph {
			ans = append(ans, i)
		}
		i++
		j++
	}
	return ans
}

func findAnagrams2(s, p string) (ans []int) {
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return
	}

	var sCount, pCount [26]int
	for i, ch := range p {
		sCount[s[i]-'a']++
		pCount[ch-'a']++
	}
	if sCount == pCount {
		ans = append(ans, 0)
	}

	for i, ch := range s[:sLen-pLen] {
		sCount[ch-'a']--
		sCount[s[i+pLen]-'a']++
		if sCount == pCount {
			ans = append(ans, i+1)
		}
	}
	return
}

func main() {
	s := "cbaebabacd"
	p := "abc"
	fmt.Println(findAnagrams(s, p))
	s = "abab"
	p = "ab"
	fmt.Println(findAnagrams(s, p))
}
