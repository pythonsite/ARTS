package main

import (
	"flag"
	"fmt"
	"os"
)

var name = flag.String("name", "everyone", "The greet object.")

func main() {
	flag.Usage = func(){
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	}
	flag.Parse()
	fmt.Printf("Hello %s!\n", *name)
}
