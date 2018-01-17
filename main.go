package main

import (
	"fmt"
	"net/http"
	"os"
)

const (
	port = ":80"
)

var calls = 0

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	calls++
	fmt.Fprintf(w, "Someone has called me in pod %s %d times.\n", os.Getenv("MY_POD_NAME"), calls)
}

func init() {
	fmt.Printf("Started server at http://localhost%v.\n", port)
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(port, nil)
}

func main() {}
