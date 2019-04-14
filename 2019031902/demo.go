package main

import (
	"flag"
	"ARTS/2019031902/lib"
	// in "ARTS/2019031902/lib/internal"  这种导入的方式是会报错的
	// "os"
)

var name string

func init() {
	flag.StringVar(&name, "name","everyone", "The greeting object.")
}

func main(){
	flag.Parse()
	lib.Hello(name)
}