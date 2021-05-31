package main

import "fmt"

func isPowerOfTwo1(n int) bool {
	if n<=0{
		return false
	}
	count:=0
	for i := 0;; i++ {
		if n==0{
			break
		}
		if n&1==1{
			count++
		}
		n=n>>1
	}
	return count==1
}

func main() {
	fmt.Println(isPowerOfTwo1(4))
}

func isPowerOfTwo2(n int) bool {
	return (n&n-1==0)&&n>0
}