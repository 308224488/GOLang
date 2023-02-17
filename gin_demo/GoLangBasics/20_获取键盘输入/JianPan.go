package main

import "fmt"

func main()  {
	/*fmt.Println() 打印并换行
	fmt.Printf()    打印格式化
	fmt.Print()     打印*/
/*	fmt.Scanln()    接收输入并换行
	fmt.Scanf()     接收格式化输入
	fmt.Scan()      接收输入*/
	var  a int
	var  b  float64
	fmt.Println("请输入两个数:1、整数2.浮点数")
	fmt.Scanln(&a,&b)
	fmt.Println("a:",a)
	fmt.Println("b:",b)
}