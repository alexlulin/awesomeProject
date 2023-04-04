package main

import (
	"fmt"
)

func main(){
	s:=make([]string,3)
	//fmt.Println("emp:",s)

	//我们可以和数组一样设置和得到值
	s[0]="a"
	s[1]="b"
	s[2]="c"
	fmt.Println("set:",s)
	fmt.Println("get:",s[2])

	//len 返回 slice 的长度
	fmt.Println("len:",len(s))

	s=append(s,"d")
	s=append(s,"e","ff")
	fmt.Println("append:",s)

	c:=make([]string,len(s))
	copy(c,s)
	fmt.Println("sl1:",c)

	l := s[2:]
	fmt.Println("sl2:",l)

	t:=[]string{"g","h","j"}
	fmt.Println("dc1:",t)


	twoD := make([][]int,3)
	for i:=0;i<3;i++{
		innerLen :=i+1
		twoD[i] = make([]int,innerLen)
		for j:=0;j<innerLen;j++{
			twoD[i][j] = i+j
		}
	}

	fmt.Println("2d:",twoD)

}
