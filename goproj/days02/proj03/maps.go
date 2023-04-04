package main

import "fmt"

func main()  {

	m:=make(map[string]int)
	m["k1"]=7
	m["k2"]=13
	fmt.Println("map:",m)

	//使用例如`fmt.Println ` 来打印一个map将会 输出所有的键值对
	//使用·name[key]· 来获取键值
	v1 := m["k1"]
	fmt.Println("v1：",v1)
	fmt.Println("len：",len(m))

	//内建的`delete` 可以从一个map 中移除键值对

	delete(m,"k2")
	fmt.Println("map:",m)

	_,prs:=m["k2"]
	fmt.Println("prs:",prs)

	n:=map[string]int{"foo":1,"bar":2}
	fmt.Println("map:",n)
}
