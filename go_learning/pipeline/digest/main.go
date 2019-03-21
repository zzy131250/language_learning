package main

import (
	"fmt"
	"os"
	"sort"
)

// 参考：https://blog.golang.org/pipelines
func main() {
	// Calculate the MD5 sum of all files under the specified directory,
	// then print the results sorted by path name.
	// m, err := MD5All(os.Args[1])
	// m, err := MD5AllParallel(os.Args[1])
	m, err := MD5AllBoundedParallel(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	var paths []string
	for path := range m {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		fmt.Printf("%x  %s\n", m[path], path)
	}
}
