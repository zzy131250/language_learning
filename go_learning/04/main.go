package main

import (
	"flag"
	"fmt"
)

func main() {
	// var name string
	// flag.StringVar(&name, "name", "everyone", "A greeting object.")
	
	// var name = *flag.String("name", "everyone", "A greeting object.")
	
	// name := *flag.String("name", "everyone", "A greeting object.")

	name := *getTheFlag()

	flag.Parse()
	fmt.Printf("Hello, %v\n", name)
}

func getTheFlag() *string {
	return flag.String("name", "everyone", "A greeting object.")
}
