
package main

import (
	"fmt"
	"net/http"
	"runtime"
	"log"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world, I'm running on %s with an %s CPU ", runtime.GOOS, runtime.GOARCH)
}

func main() {
	http.HandleFunc("/", indexHandler)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
