package main



func totalHammingDistance(nums []int) int {
	count:=0
	t:=1
	count0:=0
	for i := 0; ; i++ {
		if t==0x80000000{
			break
		}
		for j := 0; j < len(nums); j++ {
			if nums[j]&t==0{
				count0++
			}
		}
		count+=count0*(len(nums)-count0)
		count0=0
		t=t<<1
	}
	return count
}



