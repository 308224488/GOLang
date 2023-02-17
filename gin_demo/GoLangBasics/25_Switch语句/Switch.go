package main

import "fmt"

func main()  {
	var score int
	fmt.Println("请输入您的成绩，不输就把你噶了：")
	fmt.Scan(&score)
	switch score {
	case 90:
		fmt.Println("A")
	case 80:
		fmt.Println("B")
	case 70,60,50:
		fmt.Println("拉出去埋了吧")
	default:
		fmt.Println("拉出去枪毙了")
	}
}