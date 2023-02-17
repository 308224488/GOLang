package main

import "fmt"

func main()  {
	str:="DongYuRong"
	fmt.Println(str)
	fmt.Println(len(str))
    fmt.Println("打印字节:",str[2])
	fmt.Printf("打印出的结果为:%c",str[2])
	fmt.Println()
	//for循环
	for i := 0; i <len(str) ; i++ {
		fmt.Printf("%c",str[i])
	}
	fmt.Println()
	//for Rang循环  遍历数组、切片
	//返回对应下标和数组的值
	for i,v:=range str{
		fmt.Print(i)
		fmt.Printf("%c",v)
		fmt.Println()
	}

}