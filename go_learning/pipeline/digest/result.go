package main

import "crypto/md5"

type result struct {
	path string
	sum  [md5.Size]byte
	err  error
}
