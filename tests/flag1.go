package main

import (
	"flag"
	"fmt"
)

/***
# go flag 包/命令行参数

*/

var (
	confPath string
	env      string
	host     string
)

func main() {
	//flag.StringVar(&env, "e", "defult env val", "set env value")
	//// 必须调用 Parse 才能解析到对应的命令行参数
	//flag.Parse()
	//fmt.Printf("env: %s\n", env)

	tmpH := flag.String("h", "defult env val", "set env value")
	flag.Parse()
	host = *tmpH
	fmt.Printf("host: %s\n", host)
}
