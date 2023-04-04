package main

import (
	"fmt"
	"sort"
)

func main() {

	strs :=[]string{"c","a","b"}
	sort.Strings(strs)
	fmt.Println("Strings:",strs)

	// ’int‘ 排序的例子
	ints:= []int{7,8,2,3}
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Println("Ints:",ints)

	//用'sort' 来检查一个序列是不是已经排好序了
	s:= sort.IntsAreSorted(ints)
	fmt.Println("sorted：",s)
}
