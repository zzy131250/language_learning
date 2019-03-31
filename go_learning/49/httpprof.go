package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

// /debug/pprof
func main() {
	log.Println(http.ListenAndServe("localhost:8082", nil))
}
