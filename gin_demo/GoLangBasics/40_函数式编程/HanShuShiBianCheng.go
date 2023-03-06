package main

import "fmt"

func main()  {
  t1:=add(1,2)
  fmt.Println(t1)
 f:= test(2,3,add)
 fmt.Println(f)
 f2:=test(3,2,sub)
 fmt.Println(f2)
}
//高阶函数  可以将一个函数座位参数
func test(a,b int, fun func(int, int) int) int {
	r:=fun(a,b)
	return r
}
func add(a,b int) int  {
	c:=a+b
	return c
}
func sub(a,b int) int  {
	c:=a-b
	return c
}