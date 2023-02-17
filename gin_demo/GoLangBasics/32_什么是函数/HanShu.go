package main

import "fmt"

func main()  {
	var(
		a int
		b  int
	)
   fmt.Println("请输入你想输入的第一个数字:")
   fmt.Scan(&a)
	fmt.Println("请输入你想输入的第二个数字:")
	fmt.Scan(&b)
   fmt.Println("结果为:",test(a,b))

}
/*func 函数名称(参数列表)(返回值)  {
	return
}*/
func test(a int,b int)(int)  {
	c:=a*b
	return c
}
