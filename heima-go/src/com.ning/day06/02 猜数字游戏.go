package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var num int = rand.Intn(1000)
	for{
		if num >= 100 && num <= 999{
			break
		}
		num = rand.Intn(1000)
	}
	var slice = make([]int,3)
	GetNum(slice,num)
	fmt.Println(slice)

	for{
		fmt.Println("请输入三位数字:")
		var input int
		fmt.Scan(&input)
		if !(input >= 100 && input <= 999){
			fmt.Println("数字只能介于100-999之间")
			continue
		}
		var inputSlice[]int = make([]int,3)
		GetNum(inputSlice,input)
		var b bool = true
		for idx,data := range slice{
			if data > inputSlice[idx]{
				fmt.Println("第",idx + 1 ,"位数字偏小")
				b = false
				break;
			}else if data < inputSlice[idx]{
				fmt.Println("第",idx + 1,"位数字偏大")
				b = false
				break;
			}else{
				//fmt.Println("第",idx + 1,"位数字正确")
			}
		}
		if b {
			break;
		}
	}
	fmt.Println("恭喜你，数字猜对了！！！")


}
func GetNum(slice[]int,num int )  {
	slice[0] = num / 100
	slice[1] = num % 100 / 10
	slice[2] = num % 10
}
