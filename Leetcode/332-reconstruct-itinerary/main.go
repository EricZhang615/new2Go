package main

import (
	"fmt"
	"sort"
)

type target struct {
	name string
	cnt  int
}

type targets []*target

func (t targets) Len() int {
	return len(t)
}

func (t targets) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t targets) Less(i, j int) bool {
	return t[i].name < t[j].name
}

func findItinerary(tickets [][]string) []string {
	var path []string
	set := make(map[string]targets)
	var bc func() bool
	bc = func() bool {
		if len(path) == len(tickets)+1 {
			return true
		}
		for _, t := range set[path[len(path)-1]] {
			if t.cnt > 0 {
				path = append(path, t.name)
				t.cnt--
				if bc() {
					return true
				}
				path = path[:len(path)-1]
				t.cnt++
			}
		}
		return false
	}
	for _, ticket := range tickets {
		if set[ticket[0]] == nil {
			set[ticket[0]] = append(make(targets, 0), &target{ticket[1], 1})
		} else {
			for i, t := range set[ticket[0]] {
				if t.name == ticket[1] {
					t.cnt++
					break
				}
				if i == len(set[ticket[0]])-1 {
					set[ticket[0]] = append(set[ticket[0]], &target{ticket[1], 1})
					break
				}
			}
		}

	}
	for _, t := range set {
		sort.Sort(t)
	}
	path = append(path, "JFK")
	bc()
	return path
}

func main() {
	// [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
	fmt.Println(findItinerary([][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}))
}
