package main

import (
	"fmt"
)

func main()  {
	var sum int
	//计算1到10之和
	for i := 0; i <=9999 ; i++ {

		sum=sum+i
	}
	fmt.Println(sum)
}