package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for _, s := range strs {
		aph := [26]int{}
		for _, i := range s {
			aph[i-'a']++
		}
		m[aph] = append(m[aph], s)
	}
	var ans [][]string
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
	strs = []string{""}
	fmt.Println(groupAnagrams(strs))
}
