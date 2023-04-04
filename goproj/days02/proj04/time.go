package main

import (
	"fmt"
	"time"
)

func main() {
	p:=fmt.Println

	//得到当前的实际
	now :=time.Now()
	p(now)

	//通过提供的年月日等消息，你可以构建一个’time'
	//时间总是关联着位置信息，例如时区
	then :=time.Date(2023,3,31,20,34,58,65138738,time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))


	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())

}