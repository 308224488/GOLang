package main

import "fmt"
/*Go语言的结构体嵌套
1、模拟继承性:is  a
type A struct{
	field
	type B struct{
	A  匿名字段
	 b就可以直接访问a中的属性了
	2、模拟聚合关系: has - a
	type C struct{
	fieud
	type D struct{
	c C 聚合关系
	 d 就不能直接访问 c 当中的属性，而是需要通过 D.c.xx 访问*/
func main()  {
demo1()
}
func demo1()  {
	s1:=User{"礼遇",20}
	s2:=UserInfo{User{"礼遇2",25},1}//继承User结构体
	fmt.Println(s1,s2)
	var s3 UserInfo
	s3.name="礼遇3"
	s3.sex=1
	s3.age=30
	fmt.Println(s3)
}
//定义一个父类结构体
type User struct {
	name string
	age  int
}
//定义一个子类结构体
type UserInfo struct {
	User     //匿名变量  继承
	sex int  //子类自己的属性字段
}