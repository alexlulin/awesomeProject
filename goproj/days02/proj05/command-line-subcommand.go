package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	fooCmd :=flag.NewFlagSet("foo",flag.ExitOnError)
	fooEnable:=fooCmd.Bool("enable",false,"enable")
	fooName:= fooCmd.String("name","","name")

	//对于不同的子命令，我们可以定义不同的flag
	barCmd := flag.NewFlagSet("bar",flag.ExitOnError)
	barlevel := barCmd.Int("level",0,"level")

	//期望前面定义的子命令作为第一个参数传人
	if len(os.Args)<2{
		fmt.Println("expected 'foo' or 'bar' subcommnad")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println(" enable:",*fooEnable)
		fmt.Println(" name:",*fooName)
		fmt.Println(" tail:",fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println(" level:",*barlevel)
		fmt.Println(" tail:",barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)

	}
}
