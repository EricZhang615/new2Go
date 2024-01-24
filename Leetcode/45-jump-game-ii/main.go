package main

func jump(nums []int) int {
	curDis := 0
	nxtDis := 0
	res := 0
	for i := 0; i < len(nums)-1; i++ {
		nxtDis = max(nxtDis, i+nums[i])
		if i == curDis {
			curDis = nxtDis
			res++
		}
	}
	return res
}
