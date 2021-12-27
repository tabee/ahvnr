package main

import (
	"fmt"
	"net/http"

	c "github.com/tabee/ahvvalidieren"
)

func main() {

	// http://localhost/756.3903.3333.83 > true
	// http://localhost/756.3903.3333.84 > false

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		x := r.URL.Path[1:]
		v, _ := c.Validate(x)
		fmt.Fprintf(w, "You've requested %v result ist: %v", x, v)
	})

	http.ListenAndServe(":80", nil)
}
