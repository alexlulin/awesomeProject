package main

import "fmt"

func sum(nums ...int){
	fmt.Println(nums," ")
	total :=0
	for _,num := range nums{
		total += num
	}
	fmt.Println(total)
}
func main() {
	//变参函数使用常规的调用方式，传入独立的参数
	sum(1,2)
	sum(1,2,3)
	sum(1,2,3,4,5,6,7,8,9,10)
	nums := []int{1,2,3,4}
	sum(nums...)
}
