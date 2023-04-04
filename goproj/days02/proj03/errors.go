package main

import (
	"errors"
	"fmt"
)

func f1(age int) (int,error){
	if age ==42 {
		return -1,errors.New("can't work with 42")
	}

	return age +3 ,nil
}

type ageError struct {
	age int
	prob string
}

func(e * ageError) Error() string{
	return fmt.Sprintf("%d-%s",e.age, e.prob)
}

func f2(age int)(int,error){
	if age ==42{
		return -1,&ageError{age," can't work with it"}
	}
	return age+3 ,nil
}

func main() {
	for _, i :=range []int{7,42}{
		if r,e:=f1(i);e !=nil{
			fmt.Println("f1 failed:",e)
		}else {
			fmt.Println("f1 worked:",r)
		}
	}

	for _,i :=range []int{7,42}{
		if r,e:=f2(i);e !=nil{
			fmt.Println("f2 failed:",e)
		}else {
			fmt.Println("f2 worked:",r)
		}
	}

	_,e :=f2(42)

	if ae,ok := e.(*ageError);ok{
		fmt.Println(ae.age)
		fmt.Println(ae.prob)
	}
}