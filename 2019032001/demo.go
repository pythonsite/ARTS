package main

import (
	"fmt"
)

func main() {
	var name = "icode"
	value, ok := interface{}(name).(string)
	fmt.Println(value)
	fmt.Println(ok)
}