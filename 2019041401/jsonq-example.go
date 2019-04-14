package main

import (
	"fmt"
	"log"

	"github.com/thedevsaddam/gojsonq"
)

func main() {
	jq := gojsonq.New().File("./sample-data.json")
	res := jq.Find("vendor.items.[1].name")

	if jq.Error() != nil {
		log.Fatal(jq.Error())
	}

	fmt.Println(res)
}
