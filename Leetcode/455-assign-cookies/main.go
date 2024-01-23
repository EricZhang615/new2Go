package main

import "slices"

func findContentChildren(g []int, s []int) int {
	slices.Sort(g)
	slices.Sort(s)
	i := 0
	for j := 0; j < len(s) && i < len(g); j++ {
		if s[j] >= g[i] {
			i++
		}
	}
	return i
}
