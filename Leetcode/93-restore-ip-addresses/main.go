package main

import (
	"fmt"
	"strconv"
)

func isValid(s string) bool {
	if s == "" || (s[0] == '0' && len(s) != 1) {
		return false
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	if i >= 0 && i <= 255 {
		return true
	}
	return false
}

func restoreIpAddresses(s string) []string {
	var res []string
	var path []byte
	var bc func(index int, cnt int)
	bc = func(index int, cnt int) {
		if cnt == 0 {
			if isValid(s[index:]) {
				res = append(res, string(append(path, s[index:]...)))
			}
			return
		}
		for i := index; i < len(s) && i-index+1 <= 3; i++ {
			if isValid(s[index : i+1]) {
				path = append(path, s[index:i+1]...)
				path = append(path, '.')
				bc(i+1, cnt-1)
				path = path[:len(path)-(i-index+1)-1]
			} else {
				break
			}
		}
	}
	if len(s) < 4 || len(s) > 12 {
		return res
	}
	bc(0, 3)
	return res
}

func main() {
	fmt.Println(restoreIpAddresses("101023"))
}
