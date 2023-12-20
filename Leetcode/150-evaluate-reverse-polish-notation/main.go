package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	calc := map[string]int{"+": 0, "-": 0, "*": 0, "/": 0}
	var stack []string
	for i := 0; i < len(tokens); i++ {
		if _, e := calc[tokens[i]]; !e {
			stack = append(stack, tokens[i])
			continue
		}
		a, _ := strconv.Atoi(stack[len(stack)-2])
		b, _ := strconv.Atoi(stack[len(stack)-1])
		stack = stack[:len(stack)-1]
		switch tokens[i] {
		case "+":
			stack[len(stack)-1] = strconv.Itoa(a + b)
		case "-":
			stack[len(stack)-1] = strconv.Itoa(a - b)
		case "*":
			stack[len(stack)-1] = strconv.Itoa(a * b)
		case "/":
			stack[len(stack)-1] = strconv.Itoa(a / b)
		}
	}
	ans, _ := strconv.Atoi(stack[0])
	return ans
}

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}
