package main

import "fmt"

/*。无参无返回值函数
。有一个参数的函数
。有两个参数的函数
。有一个返向值的函数
。有多个返回值的函数*/
func main()  {
	var(
		a int
		b string
		c int
	)
   test_01()
   test_02(a)
   test_03(b,c)
	fmt.Println("",test_04())
	h,i:=test_05()
	fmt.Println("",h,i)
}
//无参无返回值函数
func test_01()  {
	fmt.Println("李遇")
}
//有一个参数的函数
func test_02(a int)  {
	a=666
	fmt.Println("李遇说:",a)
}
//有两个参数的函数
func test_03(b string,c int)  {
	b="李遇说:今晚吃土豆丝粉条"
	c=6666
	fmt.Println("",b,c)
}
//有一个返回值的函数
func test_04()(int){
	d:=886
	return d
}
//有一个返回值的函数
func test_05()(int,string){
	e:=886
	f:="小可爱"
	return e,f
}