package main

func partitionLabels(s string) []int {
	m := make([]int, 26)
	for i, c := range s {
		m[c-'a'] = i
	}
	res := []int{}
	cur := m[s[0]-'a']
	start := 0
	for i := 0; i < len(s); i++ {
		cur = max(m[s[i]-'a'], cur)
		if cur == i {
			res = append(res, i+1-start)
			start = i + 1
		}
	}
	return res
}
