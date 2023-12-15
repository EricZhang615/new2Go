package main

import "fmt"

func reverseWords(s string) string {
	ss := []byte(s)
	var ans []byte
	i, j := len(ss)-1, len(ss)-1
	for j >= 0 {
		if ss[j] == ' ' {
			j--
			continue
		}
		i = j
		for i >= 0 && ss[i] != ' ' {
			i--
		}
		for k := i + 1; k < j+1; k++ {
			ans = append(ans, ss[k])
		}
		ans = append(ans, ' ')
		j = i - 1
	}
	return string(ans[:len(ans)-1])
}

func reverseWords2(s string) string {
	//1.使用双指针删除冗余的空格
	slowIndex, fastIndex := 0, 0
	b := []byte(s)
	//删除头部冗余空格
	for len(b) > 0 && fastIndex < len(b) && b[fastIndex] == ' ' {
		fastIndex++
	}
	//删除单词间冗余空格
	for ; fastIndex < len(b); fastIndex++ {
		if fastIndex-1 > 0 && b[fastIndex-1] == b[fastIndex] && b[fastIndex] == ' ' {
			continue
		}
		b[slowIndex] = b[fastIndex]
		slowIndex++
	}
	//删除尾部冗余空格
	if slowIndex-1 > 0 && b[slowIndex-1] == ' ' {
		b = b[:slowIndex-1]
	} else {
		b = b[:slowIndex]
	}
	//2.反转整个字符串
	reverse(b)
	//3.反转单个单词  i单词开始位置，j单词结束位置
	i := 0
	for i < len(b) {
		j := i
		for ; j < len(b) && b[j] != ' '; j++ {
		}
		reverse(b[i:j])
		i = j
		i++
	}
	return string(b)
}

func reverse(b []byte) {
	left := 0
	right := len(b) - 1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

func main() {
	fmt.Println(reverseWords("the sky is blue"))
	fmt.Println(reverseWords("  hello world  "))
	fmt.Println(reverseWords("a good   example"))
}
