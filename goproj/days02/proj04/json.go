package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Response1 struct {
	Page int
	Fruits []string
}

type Response2 struct{
	Page int `json:"page"`
	Fruits []string `json:fruits`
}

func main() {

	/*
	首先我们来看一下基本数据类型到json 字符串的编码过程 这里一些原子值的例子
	 */

	bloB,_:=json.Marshal(true)
	fmt.Println(string(bloB))

	intB,_:=json.Marshal(1)
	fmt.Println(string(intB))

	fltB,_:=json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB,_:=json.Marshal("goland")
	fmt.Println(string(strB))

	//切片和map 编码成json 数组和对象的例子
	slcD:=[]string{"apple","peach","pear"}
	slcB,_:=json.Marshal(slcD)
	fmt.Println(string(slcB))

	// JSON包可以自动的编码你的定义类型
	//编码仅 输出可导出的字段，并且默认使用他们的名字作为JSON数据的键
	res1D := &Response1{
		Page: 1,
		Fruits: []string{"apple","peach","pear"}}
	res1B,_:=json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D:=&Response2{
		Page:1,
		Fruits:[]string{"apple","peach","pear"}}
	res2B,_:=json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt:=[]byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}

	//这里是实际的解码和相关的错误检查
	if err:= json.Unmarshal(byt,&dat);err !=nil{
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)// 转换成float64类型
	fmt.Println(num)

	//访问嵌套的值需要一些列的转化
	strs :=dat["strs"].([]interface{})
	str1:=strs[0].(string)
	fmt.Println(str1)

	str :=`{"page":1,"fruits":["apple","peach"]}`
	res:=&Response2{}
	json.Unmarshal([]byte(str),&res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc:=json.NewEncoder(os.Stdout)
	d:=map[string]int{"apple":5,"lettcue":7}
	enc.Encode(d)
}


