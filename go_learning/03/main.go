package main

import (
	"flag"
)

var name = flag.String("name", "everyone", "A greeting object.")

func main() {
	flag.Parse()
	hello(*name)
}
