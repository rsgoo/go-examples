package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	OsFlag()
}

func OsFlag() {
	//go run example_02.go -node 11313
	s := flag.String("node", "hello,world", "string help text")
	flag.Parse()
	fmt.Println("value ", *s)
}

func OsArgs() {
	for key, val := range os.Args {
		if key == 0 {
			continue
		}
		fmt.Println("argument ", key, " is ", val)
	}
}
