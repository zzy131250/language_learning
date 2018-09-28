package main

import (
	"os"
	"flag"
	"fmt"
)

var name = flag.String("name", "everyone", "A greeting object.")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	flag.Parse()
	fmt.Printf("Hello, %s\n", *name)
}
