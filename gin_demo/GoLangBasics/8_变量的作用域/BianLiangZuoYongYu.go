package main

import "fmt"

var  userName string="宝宝" //全局变量

func main()  {
	userName:="李遇"//局部变量   //GO语言中局部变量可以和全局变量使用同一变量名，但是采用就近原则
	fmt.Println(userName)
	test()
}
func test()  {
	fmt.Println(userName)
}