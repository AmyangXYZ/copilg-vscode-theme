package main

import (
	"fmt"
	"net/http"
	"time"
)

type Biu struct {
	ID int `json:"id"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
