package main

import (
	"slices"
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || a[0] == b[0] && a[1] < b[1]
	})
	var res [][]int
	for _, p := range people {
		res = slices.Insert(res, p[1], p)
	}
	return res
}
