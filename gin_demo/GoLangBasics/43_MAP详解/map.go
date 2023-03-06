package main

import (
	"fmt"
)

func main()  {
demo3()
}
func demo1() {
	//map类型的定义 [key]value  引用类型
	var map1 map[int]string
	var map2 = make(map[int]string)
	var map3 = map[string]string{"姓名": "李遇", "年龄": "二十"}
	fmt.Println(map1, map2, map3)
	fmt.Println(map1 == nil)
	fmt.Println(map2 == nil)// make创建的map是真正意义上的空map
}
func demo2()  {
	var map1 =make( map[int]string)
	map1[0]="20"
	map1[1]="李遇"
	map1[3]="666"
	fmt.Println(map1)
    fmt.Println(map1[1])
	for i := 0; i <len(map1) ; i++ {
		fmt.Println(map1[i])
	}
	//通过 ok-item判断value是否存在
	value,ok:=map1[5]
	if ok {
       fmt.Println("存在",value)
	}else{
		fmt.Println("不存在")
	}
	//修改map
	map1[0]="667666"
	fmt.Println(map1)
	//map 删除
	delete(map1,1)
	fmt.Println(map1)
	//map
	for k,v := range map1 {
		fmt.Println(k,v)
	}
}
func demo3()  {
/*需求:
	1、创建map米存储人的信息，name、age、sex、addr
	2、每个map存一个人的信息
	3、将这些map存入到 slice 中
	4、打印遍历输出*/
	user1:=map[string]string{"姓名":"李遇","年龄":"0岁","性别":"男"}
	user2:=map[string]string{"姓名":"钱多多","年龄":"1岁","性别":"女"}
	user3:=map[string]string{"姓名":"鲁迅","年龄":"10岁","性别":"男"}
	fmt.Println(user1)
	fmt.Println(user2)
	fmt.Println(user3)
	userDatas:=make([]map[string]string,0,3)
	userDatas= append(userDatas,user1,user2,user3)
	fmt.Println(userDatas)
	for _,i := range userDatas {
		fmt.Println(i["姓名"],i["年龄"])
	}
}