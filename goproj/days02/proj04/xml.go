package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct{
	XMLName xml.Name `xml:"plant"`
	Id int `xml:"id,attr"`
	Name string `xml:"name"`
	Origin []string `xml:"origin"`
}

func (p Plant)String() string{
	return fmt.Sprintf("Plant id=%v,name=%v,origin=%v",p.Id,p.Name,p.Origin)
}

func main() {
	coffee:=&Plant{Id:27,Name: "coffee"}
	coffee.Origin=[]string{"Ethiopia","Brazil"}

	out,_:=xml.MarshalIndent(coffee," "," ")
	fmt.Println(string(out))
	fmt.Printf("------------------2-------------------------\n")
	fmt.Println(xml.Header+string(out))

	fmt.Printf("-------------------3------------------------\n")
	var p Plant
	if err:= xml.Unmarshal(out,&p);err !=nil{
		panic(err)
	}

	fmt.Println(p)

	fmt.Printf("-------------------4------------------------\n")

	tomato:=&Plant{Id:81,Name: "Tomato"}
	tomato.Origin=[]string{"Mexico","California"}


	//字段标签告诉编译器嵌套 下面的plant
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants []*Plant `xml:"parent>child>plant"`
	}
	nesting := &Nesting{}
	nesting.Plants =[]*Plant{coffee,tomato}
	out,_=xml.MarshalIndent(nesting," "," ")
	fmt.Println(xml.Header+string(out))
}
