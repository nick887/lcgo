package main

func isPowerOfFour(n int) bool {
	if n<=0{
		return false
	}
	count :=0
	for i := 0; ; i++ {
		if n==0{
			break
		}
		if n&1==1{
			count++
		}
		n=n>>1
		if n&1==1{
			return false
		}
		n=n>>1
	}
	return count==1
}


func isPowerOfFour1(n int) bool {
	return n>0 && n&(n-1)==0 && n&0xaaaaaaaa==0
}
