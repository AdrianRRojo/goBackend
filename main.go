package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", getLanding)

	http.ListenAndServe(":8000", nil)
}
