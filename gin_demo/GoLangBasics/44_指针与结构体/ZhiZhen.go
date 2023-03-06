package main

import (
	"fmt"
	"gin_demo/GoLangBasics/ProjectPublicPage"
)

func main()  {
demo9()
}
//初识指针
func demo1()  {
	//& 取地址符   通过内存地址修改变量的值    *取内存地址存储的值
	a:=10
	b:=&a
	fmt.Println("变量a的值：",a)
	fmt.Println("变量a的内存地址：",&a)
	fmt.Println("变量b的内存地址：",b)
	fmt.Println("变量b的内存地址存储的值：",*b)
    *b=30
	fmt.Println("变量a的值：",a)
}
//指针的使用
func demo2()  {
	a:=10
	fmt.Println("变量a的值：",a)
	//声明一个变量指针
	var b *int
	b=&a
	fmt.Println("变量a的内存地址：",&a)
	fmt.Println("变量b的内存地址：",b)
	fmt.Println("变量b的指针的内存地址：",&b)
	fmt.Println("变量b的值：",*b)
	*b=50
	fmt.Println("变量a的值：",a)
	fmt.Println("变量b的值：",*b)
	//再声明一个变量指针
	var  c **int
	c=&b
	fmt.Println("c的内存地址：",c)
	fmt.Println("变量c的指针的内存地址：",&c)
	fmt.Println("变量c的值：",**c)

}
//指针与数组
func demo3()  {
	arr1:=[5]int{1,2,3,4,5}
	var a *[5]int
	a=&arr1
	fmt.Println("arr1的值",arr1)
	fmt.Println("a的内存地址",a)
	fmt.Println("a的指针指向的内存地址",&a)
	fmt.Println("a的值",*a)
	//古老写法
	(*a)[3]=200
	fmt.Println("arr1的值",arr1)
	fmt.Println("a的值",*a)
	//简化写法
	a[2]=300
	fmt.Println("arr1的值",arr1)
	fmt.Println("a的值",*a)
}
//指针与函数
func demo4()  {
	b:=test1()
	fmt.Println("b的内存空间:",b)
	fmt.Println("b的指针指向的内存地址:",&b)
	fmt.Println("b的值",*b)
	c:=3
	test2(&c)

}
//结构体的定义与使用
func demo5()  {
	var  user1 user
	user1.name="礼遇"
	user1.age=19
	user1.sex=0
	fmt.Println(user1)
	user2:=user{
		name: "李遇",
		age: 22,
		sex: 1,
	}
	fmt.Println(user2)
	user3:=user{"洛洛",20,3}
	fmt.Println(user3)
}
//结构体指针应用
func demo6()  {
	user1:=user{"李遇",20,1}
	user1.age=30
	fmt.Println(user1)
	var user2 *user
	user2=&user1
    user2.sex=99
    fmt.Println(*user2)
    user4:=new(user)
    user4.name="555"
    fmt.Println(*user4)
}
//匿名结构体
func demo7()  {
	s1:= struct {
		name string
		age  int
		sex  int
	}{"走走",20,1}
	fmt.Println(s1)
}
//结构体嵌套
func demo8()  {
	var  user1 =NewUser{}
	user1.name="礼遇"
	user1.age=32
	user1.sex=1
	user1.userAddress=address{
		city: "杭州",
		country: "CHINA",
	}
	fmt.Println(user1)
}
//结构体跨包调用
func demo9()  {
     user:=ProjectPublicPage.UserInfo{}
     user.Name="保持一颗好的心态"
     user.Age=25
     user.Sex=1
     fmt.Println(user)
}
func test1()*[3]int  {
	a:=[3]int{1,2,3}
	return &a
}
func test2(a *int)  {
	fmt.Println("a的内存空间:",a)
	fmt.Println("a的指针指向的内存地址:",&a)
	*a=100
	fmt.Println("a的值",*a)
}
//定义一个结构体 大写  共有别的包引入也能用   小写 私有
type user struct {
	name    string
	age     int
	sex     int
}
type NewUser struct {
	name    string
	age     int
	sex     int
	userAddress address
}
type address struct {
	city    string
	country string

}