package main

import "fmt"

func fullJustify(words []string, maxWidth int) []string {
	index, lenTmp := 0, len(words[0])
	ans := []string{}
	for i := 1; i < len(words); i++ {
		word := words[i]
		if lenTmp+len(word)+1 > maxWidth {
			// 装词
			res := maxWidth - lenTmp
			numWords := i - index
			sp1 := res / max(numWords-1, 1)
			sp2 := res % max(numWords-1, 1)
			s := ""
			for j := index; j < i-1; j++ {
				s = s + words[j] + " "
				for k := 0; k < sp1; k++ {
					s += " "
				}
				if sp2 > 0 {
					s += " "
					sp2--
				}
			}
			s += words[i-1]
			for len(s) < maxWidth {
				s += " "
			}
			ans = append(ans, s)
			index = i
			lenTmp = len(word)
		} else {
			lenTmp += len(word) + 1
		}
	}
	s := ""
	for i := index; i < len(words)-1; i++ {
		s = s + words[i] + " "
	}
	s += words[len(words)-1]
	for len(s) < maxWidth {
		s += " "
	}
	ans = append(ans, s)
	return ans
}

func main() {
	fmt.Println(fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16))
}
