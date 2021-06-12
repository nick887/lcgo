package main

func findRepeatNumber(nums []int) int {
	m:=make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if m[nums[i]]==0{
			m[nums[i]]=1
		}else {
			return nums[i]
		}
	}
	return 0
}