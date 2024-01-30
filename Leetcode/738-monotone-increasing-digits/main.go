package main

func monotoneIncreasingDigits(n int) int {
	if n == 0 {
		return 0
	}
	digits := []int{}
	for n != 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	for i := 1; i < len(digits); i++ {
		if digits[i] > digits[i-1] {
			digits[i]--
			for j := i - 1; j >= 0; j-- {
				digits[j] = 9
			}
		}
	}
	res := 0
	for i := len(digits) - 1; i >= 0; i-- {
		res *= 10
		res += digits[i]
	}
	return res
}
