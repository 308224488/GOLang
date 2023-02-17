package main

import (
	"fmt"
)

func main()  {
	var  pwd1 int
	var  pwd2 int
	pwd3:=123456
	fmt.Println("请输入密码:")
	fmt.Scan(&pwd1)
	if  pwd1==pwd3{
		fmt.Println("恭喜宁，请再次输入密码：")
		fmt.Scan(&pwd2)
		if pwd2==pwd3 {
			fmt.Println("恭喜您，登录成功！")
		}else {
			fmt.Println("第二次密码错误，拉出去埋了吧！")
		}
	}else {
		fmt.Println("直接噶了,密码错误！")
	}
}