package main

import "fmt"

func minWindow(s string, t string) string {
	i, j := 0, 0
	ans := s + "a"
	t_map := make(map[uint8]int)
	for k := 0; k < len(t); k++ {
		t_map[t[k]]++
	}
	for j < len(s) {
		if _, e := t_map[s[j]]; e {
			t_map[s[j]]--
		}

		for check(t_map) {
			if j-i+1 <= len(ans) {
				ans = s[i : j+1]
			}
			if _, e := t_map[s[i]]; e {
				t_map[s[i]]++
			}
			i++
		}
		j++
	}
	if ans == s+"a" {
		return ""
	}
	return ans
}

func check(m map[uint8]int) bool {
	for _, n := range m {
		if n > 0 {
			return false
		}
	}
	return true
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t))
	s = "a"
	t = "aa"
	fmt.Println(minWindow(s, t))
}
