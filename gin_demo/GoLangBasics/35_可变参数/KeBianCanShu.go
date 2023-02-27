package main

import "fmt"

func main()  {
	test:=countNum(1,2,3,4,5,6)
	fmt.Println(test)
}
/*...int  可变参数
如何一个函数还有其他参数，那么可变参数要放在最后
一个函数里只能有一个可变参数
*/
func countNum(nums ...int) int {
	num:=0
	for i := 0; i <len(nums) ; i++ {
		num=num+nums[i]
		fmt.Println(num)
	}
	return  num
}