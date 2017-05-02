package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {

	host, _ := os.Hostname()

	fmt.Fprintf(w, "Version is %s :::: hostname is %s \n", "v1", host)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
