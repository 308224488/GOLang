package main

import "fmt"

func main()  {
	//定义一个整数
	var age int=25
	fmt.Printf("age:%T,%d",age,age)
	//定义一个浮点类型
	var pai float64=3.1415926
	fmt.Printf("pai:%T,%.2f",pai,pai)
}