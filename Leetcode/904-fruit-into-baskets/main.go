package main

import "fmt"

func totalFruit(fruits []int) int {
	i, j := 0, 0
	f := make(map[int]int)
	for j < len(fruits) {
		for j < len(fruits) {
			if len(f) > 2 {
				if _, e := f[fruits[j]]; !e {
					f[fruits[j]] = 1
				} else {
					f[fruits[j]]++
				}
				j++
				break
			} else if len(f) == 2 {
				if _, e := f[fruits[j]]; !e {
					f[fruits[j]] = 1
					break
				} else {
					f[fruits[j]]++
				}
			} else {
				if _, e := f[fruits[j]]; !e {
					f[fruits[j]] = 1
				} else {
					f[fruits[j]]++
				}
			}
			j++
		}
		if f[fruits[i]] == 1 {
			delete(f, fruits[i])
		} else {
			f[fruits[i]]--
		}
		i++
	}
	return j - i + 1
}

func main() {
	fruits := []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
	fmt.Println(totalFruit(fruits))
	fruits = []int{1, 2, 3, 2, 2}
	fmt.Println(totalFruit(fruits))
}
