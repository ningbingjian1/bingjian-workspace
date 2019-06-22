package main

import "fmt"

func main01() {
	var s1 []int = []int{0,1,2,3,4,5,6,7,8,9}
	var s2 []int = s1[2:5]
	s2[2] = 888
	fmt.Println(s2,s1)
	var s3 []int = s2[2:5];
	s3[2] = 999
	fmt.Println(s1,s2,s3)
	fmt.Println(cap(s1),cap(s2),cap(s3))
	fmt.Println(len(s1),len(s2),len(s3))
}
