package main

import "fmt"

var dict = map[byte]string{'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}

func letterCombinations(digits string) []string {
	var res []string
	var path []byte
	var bc func(index int)
	bc = func(index int) {
		if index == len(digits) {
			res = append(res, string(path))
			return
		}
		digit := digits[index]
		for i := 0; i < len(dict[digit]); i++ {
			path = append(path, dict[digit][i])
			bc(index + 1)
			path = path[:len(path)-1]
		}
	}
	if len(digits) == 0 {
		return []string{}
	}
	bc(0)
	return res
}

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("2"))
}
