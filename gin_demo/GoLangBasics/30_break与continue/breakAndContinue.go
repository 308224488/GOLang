package main

import "fmt"

func main()  {
	for i := 0; i < 10; i++ {
		if i==5 {
          continue//跳过
		}else if i==8 {
          break//停止循环
		}
		fmt.Println(i)
	}
}