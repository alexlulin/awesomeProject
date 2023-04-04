package main

import (
	"fmt"
	"strings"
)

/*
返回目标字符串't' 出现的第一个索引位置
或者在没有匹配值时 -1.
 */

func Index(vs []string,t string)int {
	for i,v :=range vs{
		if v==t{
			return i
		}
	}
	return -1
}

/**
如果目标字符串’t' 在这个切片中则返回 ‘true’
 */
func Include(vs []string,t string)bool{
	return Index(vs,t)>=0
}
// 如果这些切片中的字符串有一个满足条‘f' 则返回’true'
func Any(vs []string,f func(string)bool)bool{
	for _,v:= range vs{
		if f(v){
			return true
		}
	}
	return false
}

func All(vs []string,f func(string) bool)bool {
	for _,v:=range vs{
		if !f(v){
			return false
		}
	}
	return true
}

//返回一个包含所有切片中满足条件‘f' 的字符串的新切片
func Filter(vs []string,f func(string) bool)[]string{
	vsf :=make([]string,0)
	for _,v:= range vs{
		if f(v){
			vsf = append(vsf,v)
		}
	}
	return vsf
}

func Map(vs []string,f func(string) string) []string{
	vsm :=make([]string,len(vs))
	for i,v:=range vs{
		vsm[i]=f(v)
	}
	return vsm
}

func main() {
	//这里试试这些组合函数
	var strs=[]string{"peach","apple","pear","plum"}
	fmt.Println(Index(strs,"pear"))
	fmt.Println(Include(strs,"grape"))
	fmt.Println(Any(strs,func(v string)bool{
		return strings.HasPrefix(v,"p")
	}))

	fmt.Println(All(strs,func(v string)bool{
		return strings.HasPrefix(v,"p")
	}))
}





