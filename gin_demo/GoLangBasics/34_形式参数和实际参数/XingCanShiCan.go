package main

import "fmt"

func main()  {
num3:= max(1,2)//实参
fmt.Println(num3)
}
func max(num1,num2 int) int {//形参
	var  result int
	if num1>num2 {
		result=num1

	}else {
		result=num2
	}
	return  result
}