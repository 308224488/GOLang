package main

import "fmt"

func main()  {
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
	defer test(4)//会被延迟到最后执行
	fmt.Println("5")
}
func test(f int)  {
	fmt.Println(f)
}