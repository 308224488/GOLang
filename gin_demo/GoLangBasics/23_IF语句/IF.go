
package main

import "fmt"

func main()  {
	var score int
	fmt.Println("请输入你这次的考试成绩：")
	fmt.Scanln(&score)
	if score>=90 {
		fmt.Println("优秀")
	}else if score>=80&&score<90 {
		fmt.Println("还凑合")
	}else if score>=70&&score<80 {
		fmt.Println("一般般")
	}else {
		fmt.Println("不及格")
	}
}