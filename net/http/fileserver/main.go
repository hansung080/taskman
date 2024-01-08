package main

import (
	"flag"
	"log"
	"net/http"
)

// Run the file server with the following flags as the example below.
// $ go run main.go --addr localhost:9001 --root /your/path
var (
	addr = flag.String("addr", ":9001", "the address of the file server")
	root = flag.String("root", "/var/www", "the root directory")
)

func main() {
	flag.Parse()
	log.Printf("The file server is running at %s\n", *addr)
	log.Fatal(http.ListenAndServe(
		*addr,
		http.FileServer(http.Dir(*root)),
	))
}
