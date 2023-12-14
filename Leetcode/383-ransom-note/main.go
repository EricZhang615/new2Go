package main

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[int32]int)
	for _, v := range magazine {
		m[v]++
	}
	for _, v := range ransomNote {
		m[v]--
		if m[v] < 0 {
			return false
		}
	}
	return true
}

func canConstruct2(ransomNote string, magazine string) bool {
	record := make([]int, 26)
	for _, v := range magazine { // 通过record数据记录 magazine里各个字符出现次数
		record[v-'a']++
	}
	for _, v := range ransomNote { // 遍历ransomNote，在record里对应的字符个数做--操作
		record[v-'a']--
		if record[v-'a'] < 0 { // 如果小于零说明ransomNote里出现的字符，magazine没有
			return false
		}
	}
	return true
}
